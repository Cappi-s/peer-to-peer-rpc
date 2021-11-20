package main

import (
	"log"
	"sync"

	rpcclient "github.com/Cappi-s/peer-to-peer-rpc/client"
	"github.com/Cappi-s/peer-to-peer-rpc/server"
	"github.com/Cappi-s/peer-to-peer-rpc/service/hello"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	server := server.Server{Host: "localhost:1234", Wg: &wg}
	go server.StartServer()

	client := rpcclient.NewClient()
	reply, err := client.XmlRpcCall("HelloService.Say", &hello.Payload{
		Sender: "Pedro",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", reply.Message)

	wg.Wait()
}
