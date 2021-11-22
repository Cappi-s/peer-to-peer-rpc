package client

import (
	"encoding/json"
	"log"

	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
	"github.com/Cappi-s/peer-to-peer-rpc/util"
)

type Client struct {
	MyAddress  string
	MyNickname string
	Peers      map[string]string
}

func NewClient(myAddress string, myNickName string, peers map[string]string) *Client {
	return &Client{
		MyAddress:  myAddress,
		MyNickname: myNickName,
		Peers:      peers,
	}
}

func (c *Client) SendMessage(recipient string, content string) {
	alreadyContacted := make(map[string]string)
	alreadyContacted[c.MyNickname] = c.MyAddress

	res, err := json.Marshal(alreadyContacted)
	if err != nil {
		log.Fatal(err)
	}

	payload := chat.Payload{
		Sender:           c.MyNickname,
		Recipient:        recipient,
		Content:          content,
		AlreadyContacted: string(res),
	}

	for nick, hostAddress := range c.Peers {
		if nick == payload.Recipient {
			go util.XmlRpcCall(hostAddress, "ChatService.SetMessage", &payload)
			return
		}
	}

	for _, hostAddress := range c.Peers {
		go util.XmlRpcCall(hostAddress, "ChatService.SetMessage", &payload)
	}
}
