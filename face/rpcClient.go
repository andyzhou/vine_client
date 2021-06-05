package face

import (
	"errors"
	"fmt"
	"github.com/andyzhou/vine_client/define"
	"log"
	"net/rpc"
	"sync"
	"time"
)

/*
 * face of rpc client, implement of `IRpcClient`
 * - auto connect assigned address
 */

//inter macro define
const (
	ErrMsgOfShutDown = "connection is shut down"
	NewServerChanSize = 128
)

//one client info
type rpcClient struct {
	client *rpc.Client
	isActive bool
	upTime int64
}

//face info
type RpcClient struct {
	servers map[define.ServerAddress]*rpcClient
	newServerChan chan define.ServerAddress
	closeChan chan bool
	sync.RWMutex
}

//construct
func NewRpcClient() *RpcClient {
	//self init
	self := &RpcClient{
		servers: make(map[define.ServerAddress]*rpcClient),
		newServerChan: make(chan define.ServerAddress, NewServerChanSize),
		closeChan:make(chan bool, 1),
	}
	//spawn tick handle
	go self.tickHandle()
	return self
}

//quit
func (f *RpcClient) Quit() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("RpcClient:Quit panic, err:", err)
		}
	}()
	f.closeChan <- true
}

//server node down notify
func (f *RpcClient) NodeDown(node define.ServerAddress) bool {
	//check
	if node == "" || f.servers == nil {
		return false
	}

	//mark relate client status
	v, ok := f.servers[node]
	if !ok {
		return false
	}
	f.Lock()
	defer f.Unlock()
	v.isActive = false
	return true
}

//call batch
func (f *RpcClient) CallAll(
			rpcName string,
			args interface{},
		) error {
	var (
		errList string
		errChan = make(chan error)
	)

	//check
	if rpcName == "" || f.servers == nil || len(f.servers) <= 0 {
		return errors.New("invalid parameter")
	}

	//call on all
	for _, client := range f.servers {
		go func(client *rpc.Client) {
			err := client.Call(rpcName, args, nil)
			errChan <- err
		}(client.client)
	}

	//receive error
	for range f.servers {
		if err := <- errChan; err != nil {
			errList += err.Error() + ";"
		}
	}
	if errList == "" {
		return nil
	}
	return fmt.Errorf(errList)
}

//call method
func (f *RpcClient) Call(
			rpcName string,
			args, reply interface{},
		) error {
	//check
	if rpcName == "" {
		return errors.New("invalid parameter")
	}

	//get rpc client
	rpcClient := f.getClient()
	if rpcClient == nil {
		return errors.New("can't get rpc client")
	}

	//check client is active
	if !rpcClient.isActive {
		return errors.New("rpc client has down")
	}

	//rpc call
	err := rpcClient.client.Call(rpcName, args, reply)
	if err != nil {
		//if server shut down
		if err.Error() == ErrMsgOfShutDown {
			rpcClient.isActive = false
		}
		log.Printf("rpc call name:%v, args:%v failed, err:%v\n",
					rpcName, args, err.Error())
	}
	return err
}

//add rpc server nodes
func (f *RpcClient) AddNodes(addresses ...define.ServerAddress) bool {
	//check
	if addresses == nil || len(addresses) <= 0 {
		return false
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("RpcClient:AddNode panic, err:", err)
		}
	}()

	//loop check
	for _, address := range addresses {
		//check
		_, ok := f.servers[address]
		if ok {
			continue
		}
		//add new server
		f.newServerChan <- address
	}

	return true
}

///////////////
//private func
///////////////

//tick handle
func (f *RpcClient) tickHandle() {
	var (
		clientCheckTicker = time.NewTicker(define.RpcClientCheckTick)
		newServer define.ServerAddress
		isOk bool
	)

	defer func() {
		if f.servers != nil {
			//close rpc clients
			for _, client := range f.servers {
				if client.client != nil {
					client.client.Close()
				}
			}
			f.servers = make(map[define.ServerAddress]*rpcClient)
		}
		clientCheckTicker.Stop()
		close(f.closeChan)
	}()

	//loop
	for {
		select {
		case newServer, isOk = <- f.newServerChan:
			if isOk {
				//init new server
				f.initNewServer(newServer)
			}
		case <- clientCheckTicker.C://check active client
			f.reConnClient()
		case <- f.closeChan:
			return
		}
	}
}

//call on

//re-conn un-active client
func (f *RpcClient) reConnClient() bool {
	//check
	if f.servers == nil || len(f.servers) <= 0 {
		return false
	}
	//loop
	for nodeAddr, client := range f.servers {
		if client.isActive {
			continue
		}
		//re connect server
		f.dialServer(nodeAddr)
	}
	return true
}

//get random client
func (f *RpcClient) getClient() *rpcClient {
	//check
	if f.servers == nil {
		return nil
	}
	//try hit in cache
	for _, client := range f.servers {
		if client != nil && client.isActive {
			return client
		}
	}
	return nil
}

//get client by address
func (f *RpcClient) getClientByAddr(addr define.ServerAddress) *rpcClient {
	//check
	if f.servers == nil {
		return nil
	}
	client, ok := f.servers[addr]
	if !ok {
		return nil
	}
	return client
}

//dial server
func (f *RpcClient) dialServer(address define.ServerAddress) *rpcClient {
	//try dial server
	client, err := rpc.Dial("tcp", string(address))
	if err != nil {
		log.Printf("rpc dial %v failed, err:%v\n", address, err.Error())
		return nil
	}
	//sync into cache with locker
	rpcClient := &rpcClient{
		client: client,
		isActive:true,
		upTime: time.Now().Unix(),
	}
	f.Lock()
	defer f.Unlock()
	f.servers[address] = rpcClient
	return rpcClient
}

//new server init
func (f *RpcClient) initNewServer(node define.ServerAddress) bool {
	//check
	if node == "" {
		return false
	}

	//dial
	rpcClient := f.dialServer(node)
	if rpcClient == nil {
		return false
	}
	return true
}