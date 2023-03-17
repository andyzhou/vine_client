package main

import (
	"fmt"
	vine "github.com/andyzhou/vine_client"
	"io/ioutil"
	"log"
	"os"
)

//server define
const (
	VineMasterNode = "127.0.0.1:7777"
	FileShortUrl = "4ARF58"
)

const (
	TxtFileName = "test.txt"
	TxtFilePath = "files/test.txt"
)

//main
func main() {
	//init client
	client := vine.NewClient()

	//add master node
	err := client.AddNodes(VineMasterNode)
	if err != nil {
		log.Printf("connect nodes failed, err:%v", err.Error())
		return
	}

	//write file
	//WriteFile(client)

	//read file
	//ReadFile(client)

	//list file
	//ListFile(client)

	//list nodes
	ListChunkNodes(client)
}

//chunk node list
func ListChunkNodes(client *vine.Client)  {
	page := 1
	pageSize := 10
	nodeList, err := client.ListChunkNodes(page, pageSize)
	if err != nil {
		log.Printf("list nodes failed, err:%v", err.Error())
		return
	}
	if nodeList == nil {
		log.Printf("list node, no records")
		return
	}
	for _, v := range nodeList.Recs {
		log.Printf("v:%v\n", v)
	}
}

//delete file
func DeleteFile(client *vine.Client)  {
	//delete file data by short url
	err := client.DelFile(FileShortUrl, "")
	if err != nil {
		log.Printf("delete file failed, err:%v", err.Error())
		return
	}

	log.Printf("delete file succeed")
}

//list file
func ListFile(client *vine.Client)  {
	//get file list
	page := 1
	pageSize := 10
	fileList, err := client.ListFile(page, pageSize)
	if err != nil {
		log.Printf("list file failed, err:%v", err.Error())
		return
	}
	if fileList == nil {
		log.Printf("list file, no records")
		return
	}

	log.Printf("list file succeed, files:%v", fileList.RecSize)
	for _, v := range fileList.Files {
		log.Printf("shortUrl:%v, fileName:%v",
			v.ShortUrl, v.Name)
	}
}

//read file
func ReadFile(client *vine.Client)  {
	//read file data by short url
	fileData, err := client.ReadFile(FileShortUrl)
	if err != nil {
		log.Printf("read file failed, err:%v", err.Error())
		return
	}
	fileInfo := string(fileData)
	log.Printf("read file succeed, data:%v\n", fileInfo)
}

//write file
func WriteFile(client *vine.Client)  {
	//read file data
	fileData, err := ReadOriginFile(TxtFilePath)
	if err != nil {
		log.Printf("read origin file failed, err:%v", err.Error())
		return
	}

	//send file data to vine master
	fileShortUrl, token, err := client.WriteFile(TxtFileName, fileData)
	if err != nil {
		log.Printf("write file failed, err:%v", err.Error())
		return
	}

	log.Printf("write file succeed, short url:%v, token:%v", fileShortUrl, token)
}

//read origin file
func ReadOriginFile(filePath string) ([]byte, error)  {
	curDir, _ := os.Getwd()
	fileRealPath := fmt.Sprintf("%v/%v", curDir, filePath)
	fileData, err := ioutil.ReadFile(fileRealPath)
	return fileData, err
}