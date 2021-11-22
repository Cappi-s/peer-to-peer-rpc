package util

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
)

func XmlRpcCall(hostAddress string, method string, payload interface{}) {
	buf, _ := xml.EncodeClientRequest(method, payload)

	resp, err := http.Post(hostAddress+"/RPC", "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Printf("error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
}
