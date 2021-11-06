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

//list file
func (c *Client) ListFile(
					page,
					pageSize int,
				) (*comm.ListFileReply, error) {
	//send rpc call on master node
	args := comm.ListFileArg{
		Page: page,
		PageSize: pageSize,
	}
	reply := &comm.ListFileReply{}
	err := c.rpcClient.Call(
		fmt.Sprintf(define.MaterRpcNamePara, "PRCListFile"),
		args,
		reply,
	)
	return reply, err
}

//delete file
func (c *Client) DelFile(
					shortUrl,
					token string,
				) error {
	//check
	if shortUrl == "" || token == "" {
		return errors.New("invalid parameter")
	}

	//send rpc call on master node
	args := comm.DelFileArg{
		ShortUrl: define.ShortUrl(shortUrl),
		Token: token,
	}
	reply := comm.DelFileReply{}
	err := c.rpcClient.Call(
		fmt.Sprintf(define.MaterRpcNamePara, "RPCDelFile"),
		args,
		&reply,
	)
	return err
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
					shortUrls ... string,
				) (string, error) {
	//check
	if fileName == "" || fileType == "" || data == nil {
		return "", errors.New("invalid parameter")
	}

	//check assigned short url
	assignedShortUrl := ""
	if shortUrls != nil && len(shortUrls) > 0 {
		assignedShortUrl = shortUrls[0]
	}

	//send rpc call on master node
	args := comm.WriteFileArg{
		File:define.FileName(fileName),
		Type:define.FileType(fileType),
		Data:data,
		ShortUrl: define.ShortUrl(assignedShortUrl),
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
