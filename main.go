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
		"Diana": "http://aeb1-5-62-49-126.ngrok.io",
		"Ian":   "http://1ca6-200-196-135-247.ngrok.io",
	}

	const (
		HOST     = "http://e509-2804-56c-ffc1-c200-b1d8-ea68-f0e6-35bd.ngrok.io"
		NICKNAME = "Pedro"
	)

	server := server.Server{
		Host:       HOST,
		Wg:         &wg,
		Peers:      peers,
		MyNickname: NICKNAME,
	}
	go server.StartServer()

	client := rpcclient.NewClient(HOST, NICKNAME, peers)
	client.SendMessage("ChatService.SetMessage", &chat.Payload{
		Sender:    NICKNAME,
		Recipient: "Ian",
		Content:   "Oi, ian, tudo bem?",
	})
	client.SendMessage("ChatService.SetMessage", &chat.Payload{
		Sender:    NICKNAME,
		Recipient: "Diana",
		Content:   "Oi, diana, tudo bem?",
	})

	wg.Wait()
}
