package vine

import (
	"errors"
	"fmt"
	"github.com/andyzhou/vine_client/comm"
	"github.com/andyzhou/vine_client/define"
	"github.com/andyzhou/vine_client/face"
	"github.com/andyzhou/vine_client/iface"
)

/*
 * client face
 *
 * - opt on master node pass rpc
 */

//face info
type Client struct {
	rpcClient iface.IRpcClient
}

//construct
func NewClient() *Client {
	//self init
	self := &Client{
		rpcClient: face.NewRpcClient(),
	}
	return self
}

//quit
func (c *Client) Quit() {
	if c.rpcClient == nil {
		return
	}
	c.rpcClient.Quit()
}

//read file
func (c *Client) ReadFile(
					shortUrl string,
					offset,
					length int64,
				) ([]byte, error) {
	//check
	if shortUrl == "" {
		return nil, errors.New("invalid parameter")
	}

	//send rpc call on master node
	args := comm.ReadFileArg{
		ShortUrl:define.ShortUrl(shortUrl),
		Offset:define.Offset(offset),
		Length:define.Length(length),
	}
	reply := comm.ReadFileReply{}
	err := c.rpcClient.Call(
		fmt.Sprintf(define.MaterRpcNamePara, "RPCReadFile"),
		args,
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return reply.Data, nil
}

//write file
func (c *Client) WriteFile(
					fileName,
					fileType string,
					data []byte,
				) (string, error) {
	//check
	if fileName == "" || fileType == "" || data == nil {
		return "", errors.New("invalid parameter")
	}

	//send rpc call on master node
	args := comm.WriteFileArg{
		File:define.FileName(fileName),
		Type:define.FileType(fileType),
		Data:data,
	}
	reply := comm.WriteFileReply{}
	err := c.rpcClient.Call(
		fmt.Sprintf(define.MaterRpcNamePara, "RPCWriteFile"),
		args,
		&reply,
	)
	if err != nil {
		return "", err
	}
	return string(reply.ShortUrl), nil
}

//add master node
func (c *Client) AddNodes(nodes ...string) bool {
	//check
	if nodes == nil || len(nodes) <= 0 {
		return false
	}
	//add nodes
	serverNodes := make([]define.ServerAddress, 0)
	for _, node := range nodes {
		serverNodes = append(serverNodes, define.ServerAddress(node))
	}
	c.rpcClient.AddNodes(serverNodes...)
	return true
}
