package main

import (
	"sync"

	rpcclient "github.com/Cappi-s/peer-to-peer-rpc/client"
	"github.com/Cappi-s/peer-to-peer-rpc/server"
	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	peers := map[string]string{
		"Diana": "localhosta:1234",
		"Ian":   "localhosta:1234",
	}
	server := server.Server{
		Host:       "localhost:1234",
		Wg:         &wg,
		Peers:      peers,
		MyNickname: "Pedro",
	}
	go server.StartServer()

	client := rpcclient.NewClient(peers)
	client.SendMessage("ChatService.SetMessage", &chat.Payload{
		Sender:    "Pedro",
		Recipient: "Pedro",
		Content:   "Oi, blz?",
	})

	wg.Wait()
}
