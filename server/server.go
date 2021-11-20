package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"

	"github.com/Cappi-s/peer-to-peer-rpc/service/hello"
)

type Server struct {
	Host string
	Wg   *sync.WaitGroup
}

func (s *Server) StartServer() {

	RPC := rpc.NewServer()
	xmlrpcCoded := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCoded, "text/xml")
	RPC.RegisterService(new(hello.HelloService), "")

	http.Handle("/RPC", RPC)

	log.Printf("Server: Starting XML-RPC server on %s/RPC", s.Host)
	log.Println(http.ListenAndServe(s.Host, nil))
	s.Wg.Done()
}
