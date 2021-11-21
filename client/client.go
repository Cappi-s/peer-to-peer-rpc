package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"

	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
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

func (c *Client) SendMessage(method string, payload *chat.Payload) {
	payload.AlreadyContacted = make(map[string]string)
	payload.AlreadyContacted[c.MyNickname] = c.MyAddress

	for nick, hostAddress := range c.Peers {
		if nick == payload.Recipient {
			go c.XmlRpcCall(hostAddress, method, payload)
			return
		}
	}

	for _, hostAddress := range c.Peers {
		go c.XmlRpcCall(hostAddress, method, payload)
	}
}

func (c *Client) XmlRpcCall(hostAddress string, method string, payload *chat.Payload) {
	buf, _ := xml.EncodeClientRequest(method, payload)
	resp, err := http.Post(hostAddress+"/RPC", "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Printf("error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	var response chat.Response

	err = xml.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		fmt.Printf("error decoding response: %v", err)
	}
}
