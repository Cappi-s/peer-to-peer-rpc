package hello

import (
	"fmt"
	"log"
	"net/http"
)

type HelloService struct{}

type Payload struct {
	Sender  string
	Content string
}

type Response struct {
	Message string
}

func (h *HelloService) Say(r *http.Request, payload *Payload, response *Response) error {

	log.Println("Server: Request received from:", payload.Sender)
	response.Message = fmt.Sprintf("Hello, %s!", payload.Sender)
	return nil
}
