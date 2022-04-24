package main

import (
	vine "github.com/andyzhou/vine_client"
	"log"
)

/*
 * example of write file
 */

//inter macro define
const (
	TxtFileName = "test.txt"
	TxtFileType = "content/text"
	TxtFilePath = "files/test.txt"
)

func main() {
	//init client
	client := vine.NewClient()

	//add master node
	err := client.AddNodes(define.VineMasterNode)
	if err != nil {
		log.Printf("connect nodes failed, err:%v", err.Error())
		return
	}

	//read file data
	base := define.NewBase()
	fileData, err := base.ReadFile(TxtFilePath)
	if err != nil {
		log.Printf("read origin file failed, err:%v", err.Error())
		return
	}

	//send file data to vine master
	fileShortUrl, err := client.WriteFile(TxtFileName, fileData)
	if err != nil {
		log.Printf("write file failed, err:%v", err.Error())
		return
	}

	log.Printf("write file succeed, short url:%v", fileShortUrl)
}
