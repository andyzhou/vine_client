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
