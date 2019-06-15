package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	host       = flag.String("host", "localhost:7070", "Host")
	connNumber = flag.Int("conn", 1, "Number of connections")
)

func main() {
	flag.Parse()

	var conns []net.Conn
	for i := 0; i < *connNumber; i++ {
		c, err := net.DialTimeout("tcp", *host, 10*time.Second)
		if err != nil {
			log.Printf("failed to connect %d: %v", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(2 * time.Microsecond)
	}
	defer func() {
		for _, c := range conns {
			err := c.Close()
			if err != nil {
				log.Printf("can't close connection: %v", err)
			}
		}
	}()

	log.Printf("created %d connections", len(conns))

	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(1 * time.Second)
			conn := conns[i]
			_, err := conn.Write([]byte(fmt.Sprintf("Send from client %d\r\n", i)))
			if err != nil {
				log.Printf("can't send data from client %d: %v", i, err)
			}
		}
	}

}
