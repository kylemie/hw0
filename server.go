package main

import (
	"github/kylemie/tcp-server-client/protocol"
	"log"
	"net"
)

func serveConnections(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(protocol.ResponseOK))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
		conn.Close()
	}
}

func main() {
	l, err := net.Listen("tcp", protocol.Port)

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	serveConnections(l)
}
