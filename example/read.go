package main

import (
	vine "github.com/andyzhou/vine_client"
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
	fileData, err := client.ReadFile(define.FileShortUrl)
	if err != nil {
		log.Printf("read file failed, err:%v", err.Error())
		return
	}

	log.Printf("read file succeed, data:%v", string(fileData))
}