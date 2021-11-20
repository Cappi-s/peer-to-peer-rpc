package client

import (
	"bytes"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"

	"github.com/Cappi-s/peer-to-peer-rpc/service/hello"
)

type Client struct{}

const host = "http://localhost:1234/RPC"

func NewClient() *Client {
	return &Client{}
}

func (c *Client) XmlRpcCall(method string, payload *hello.Payload) (*hello.Response, error) {
	buf, _ := xml.EncodeClientRequest(method, payload)

	resp, err := http.Post(host, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response hello.Response
	err = xml.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
