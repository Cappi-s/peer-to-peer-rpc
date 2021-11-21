package chat

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
)

type ChatService struct {
	Host       string
	Peers      map[string]string
	MyNickname string
}

type Payload struct {
	Sender           string
	Recipient        string
	Content          string
	AlreadyContacted map[string]string
}

type Response struct{}

func (c *ChatService) SetMessage(r *http.Request, payload *Payload, response *Response) error {

	fmt.Println("Server: Request received from:", payload.Sender)

	if payload.AlreadyContacted == nil {
		payload.AlreadyContacted = make(map[string]string)
	}

	// Adiciona peers que não temos como contato na lista
	for nick, address := range payload.AlreadyContacted {
		c.Peers[nick] = address
	}

	// Sou o destinatário, exibe a mensangem
	if payload.Recipient == c.MyNickname {
		fmt.Printf("%s: %s\n", payload.Sender, payload.Content)
		return nil
	}

	// Se adiciona na lista de pessoas que já receberam esta mensagem
	payload.AlreadyContacted[c.MyNickname] = c.Host

	// Se temos o destinatário na lista de contatos, enviamos a mensagem para ela
	hostAddress, ok := c.Peers[payload.Recipient]
	if ok {
		_, err := c.XmlRpcCall(hostAddress, "ChatService.SetMessage", payload)
		return err
	}

	//  Propaga a mensagem para todos os contatos conhecidos por mim que ainda não receberam esta mensagem
	for peerNick, peerAddress := range c.Peers {
		if _, ok := payload.AlreadyContacted[peerNick]; ok {
			continue
		}

		_, err := c.XmlRpcCall(peerAddress, "ChatService.SetMessage", payload)
		return err
	}

	return nil
}

func (c *ChatService) XmlRpcCall(host string, method string, payload *Payload) (*Response, error) {
	buf, _ := xml.EncodeClientRequest(method, payload)

	resp, err := http.Post(host, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var response Response
	err = xml.DecodeClientResponse(resp.Body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &response, nil
}
