package main

import (
	"log"
	"sync"

	rpcclient "github.com/Cappi-s/peer-to-peer-rpc/client"
	"github.com/Cappi-s/peer-to-peer-rpc/server"
	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	peers := map[string]string{
		"Pedro": "localhost:1234",
		"Ian":   "localhost:1234",
	}
	server := server.Server{
		Host:       "localhost:1234",
		Wg:         &wg,
		Peers:      peers,
		MyNickname: "Diana",
	}
	go server.StartServer()

	client := rpcclient.NewClient()
	reply, err := client.XmlRpcCall("ChatService.SetMessage", &chat.Payload{
		Sender:  "Diana",
		Content: "Oi, blz?",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", reply.Message)

	wg.Wait()
}
