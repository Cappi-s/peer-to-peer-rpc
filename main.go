package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	rpcclient "github.com/Cappi-s/peer-to-peer-rpc/client"
	"github.com/Cappi-s/peer-to-peer-rpc/server"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	bytes, err := ioutil.ReadFile("peers")
	if err != nil {
		log.Fatal(err)
	}

	var host, nickname string

	peerList := strings.Split(string(bytes), "\n")
	peers := make(map[string]string)
	for i, peerString := range peerList {
		clean := strings.Split(peerString, "\r")[0]
		splited := strings.Split(clean, " ")
		nick := splited[0]
		address := splited[1]

		if i == 0 {
			host = address
			nickname = nick
			continue
		}

		peers[nick] = address
	}

	server := server.Server{
		Host:       host,
		Wg:         &wg,
		Peers:      peers,
		MyNickname: nickname,
	}
	go server.StartServer()
	wg.Wait()

	client := rpcclient.NewClient(host, nickname, peers)

	fmt.Println("Format: RecipientName Message")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Send Message: ")
		scanner.Scan()
		text := scanner.Text()
		splited := strings.Split(text, " ")
		nick := splited[0]
		message := strings.Join(splited[1:], " ")

		client.SendMessage(nick, message)
	}
}
