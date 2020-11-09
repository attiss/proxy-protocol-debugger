package main

import (
	"log"
	"net"

	proxyproto "github.com/pires/go-proxyproto"
)

func main() {
	addr := ":8080"
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not listen on %s: %q\n", addr, err.Error())
	}

	proxyListener := &proxyproto.Listener{Listener: list}
	defer proxyListener.Close()

	for {
		conn, _ := proxyListener.Accept()

		if conn.LocalAddr() == nil {
			log.Fatal("could not retrieve local address")
		}
		log.Printf("local address: %s", conn.LocalAddr().String())

		if conn.RemoteAddr() == nil {
			log.Fatal("could not retrieve remote address")
		}
		log.Printf("remote address: %q", conn.RemoteAddr().String())

		if err := conn.Close(); err != nil {
			log.Fatalf("could not close connection: %s", err.Error())
		}
	}
}
