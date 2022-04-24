package main

import (
	vine "github.com/andyzhou/vine_client"
	"github.com/andyzhou/vine_client/example/define"
	"log"
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

	//read file data by short url
	err = client.DelFile(define.FileShortUrl, "")
	if err != nil {
		log.Printf("delete file failed, err:%v", err.Error())
		return
	}

	log.Printf("delete file succeed")
}
