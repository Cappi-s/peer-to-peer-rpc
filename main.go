package main

import (
	"sync"
	"time"

	"github.com/Cappi-s/peer-to-peer-rpc/server"
	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"

	rpcclient "github.com/Cappi-s/peer-to-peer-rpc/client"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	peers := map[string]string{
		"Diana": "aeb1-5-62-49-126.ngrok.io",
		"Ian":   "1ca6-200-196-135-247.ngrok.io",
	}

	const (
		HOST     = "http://localhost:1234"
		NICKNAME = "Pedro"
	)

	server := server.Server{
		Host:       HOST,
		Wg:         &wg,
		Peers:      peers,
		MyNickname: NICKNAME,
	}
	go server.StartServer()

	time.Sleep(time.Second * 5)

	client := rpcclient.NewClient(HOST, NICKNAME, peers)
	client.SendMessage("ChatService.SetMessage", &chat.Payload{
		Sender:    NICKNAME,
		Recipient: "Diana",
		Content:   "Oi, didu, tudo bem?",
	})

	wg.Wait()
}
