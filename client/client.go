package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"

	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
)

type Client struct{}

const host = "http://localhost:1234/RPC"

func NewClient() *Client {
	return &Client{}
}

func (c *Client) XmlRpcCall(method string, payload *chat.Payload) (*chat.Response, error) {
	buf, _ := xml.EncodeClientRequest(method, payload)

	resp, err := http.Post(host, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var response chat.Response
	err = xml.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &response, nil
}
