package iface

import (
	"github.com/andyzhou/vine_client/define"
)

/*
 * interface of rpc service/client
 */

type IRpcClient interface {
	Quit()
	NodeDown(node define.ServerAddress) bool
	CallAll(rpcName string, args interface{}) error
	Call(rpcName string, args, reply interface{}) error
	AddNodes(addresses ...define.ServerAddress) bool
}