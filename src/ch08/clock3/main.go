// Demo multi-threaded clock server
// Supports a 'port' command-line flag to pass a port
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port upon which time should be published")
	flag.Parse()
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. the client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. the connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
