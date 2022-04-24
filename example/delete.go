package main

import (
	vine "github.com/andyzhou/vine_client"
	"log"
)

func main() {
	//init client
	client := vine.NewClient()

	//add master node
	err := client.AddNodes(VineMasterNode)
	if err != nil {
		log.Printf("connect nodes failed, err:%v", err.Error())
		return
	}

	//read file data by short url
	err = client.DelFile(FileShortUrl, "")
	if err != nil {
		log.Printf("delete file failed, err:%v", err.Error())
		return
	}

	log.Printf("delete file succeed")
}
