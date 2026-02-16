package main

import (
	"bufio"
	"github/kylemie/tcp-server-client/protocol"
	"log"
	"net"
)

const (
	serverHost = "localhost"
)

func main() {
	conn, err := net.Dial("tcp", serverHost+protocol.Port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	if response != protocol.ResponseOK {
		log.Fatal("not OK")
	}
}
