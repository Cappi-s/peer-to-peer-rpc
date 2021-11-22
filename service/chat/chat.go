package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Cappi-s/peer-to-peer-rpc/util"
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
	AlreadyContacted string
}

type Response struct{}

func (c *ChatService) SetMessage(r *http.Request, payload *Payload, response *Response) error {
	alreadyContacted := make(map[string]string)
	err := json.Unmarshal([]byte(payload.AlreadyContacted), &alreadyContacted)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Adiciona peers que não temos como contato na lista
	for nick, address := range alreadyContacted {
		if nick == c.MyNickname {
			continue
		}
		c.Peers[nick] = address
	}

	// Sou o destinatário, exibe a mensangem
	if payload.Recipient == c.MyNickname {
		fmt.Printf("\n%s: %s\n\n", payload.Sender, payload.Content)
		return nil
	}

	// Se adiciona na lista de pessoas que já receberam esta mensagem
	alreadyContacted[c.MyNickname] = c.Host

	// Se temos o destinatário na lista de contatos, enviamos a mensagem para ela
	hostAddress, ok := c.Peers[payload.Recipient]
	if ok {
		util.XmlRpcCall(hostAddress, "ChatService.SetMessage", payload)
		return nil
	}

	//  Propaga a mensagem para todos os contatos conhecidos por mim que ainda não receberam esta mensagem
	for peerNick, peerAddress := range c.Peers {
		if _, ok := alreadyContacted[peerNick]; ok {
			continue
		}

		res, err := json.Marshal(alreadyContacted)
		if err != nil {
			log.Fatal(err)
			return err
		}

		payload.AlreadyContacted = string(res)

		util.XmlRpcCall(peerAddress, "ChatService.SetMessage", payload)
	}

	return nil
}
