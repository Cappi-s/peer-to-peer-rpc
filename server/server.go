package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"

	"github.com/Cappi-s/peer-to-peer-rpc/service/chat"
)

type Server struct {
	Host       string
	Wg         *sync.WaitGroup
	Peers      map[string]string
	MyNickname string
}

func (s *Server) StartServer() {

	RPC := rpc.NewServer()
	xmlrpcCoded := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCoded, "text/xml")

	chatService := chat.ChatService{
		Host:       s.Host,
		Peers:      s.Peers,
		MyNickname: s.MyNickname,
	}
	RPC.RegisterService(&chatService, "")

	http.Handle("/RPC", RPC)

	log.Printf("Server: Starting XML-RPC server on %s/RPC", s.Host)
	s.Wg.Done()
	log.Println(http.ListenAndServe("localhost:1234", nil))
}
