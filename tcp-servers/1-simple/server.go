package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Panic(err)
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	var conns []net.Conn
	defer func() {
		for _, c := range conns {
			err := c.Close()
			if err != nil {
				log.Printf("can't close connection: %v", err)
			}
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temporary err: %v", ne)
				continue
			}
			log.Printf("accept err: %v", err)
			return
		}

		go handleConn(conn)

		conns = append(conns, conn)
		//if len(conns) % 100 == 0 {
		log.Printf("number of connections: %v", len(conns))
		//}
	}
}

func handleConn(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Printf("can't copy payload from connection: %v", err)
	}
}
