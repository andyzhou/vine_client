# vine_client
vine client api

# about vine
vine is a file storage service, git url:https://github.com/andyzhou/vine

# example
```
 import "github.com/andyzhou/vine_client"

 //init client
 client := vine.NewClient()
 
 //add master node
 client.AddNodes("localhost:777")
 
 //for write
 shortUrl, err := client.WriteFile(
 					owner int64,
 					fileName,
 					fileType string,
 					data []byte,
 				)
 				
 //for read
data, err := client.ReadFile(
                  shortUrl string,
                  offset,
                  length int64,
                 )

//quit
client.Quit()
```