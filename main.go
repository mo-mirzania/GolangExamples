package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		log.Panic()
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var i int
	for scanner.Scan() {
		if i == 0 {
			words := strings.Fields(scanner.Text())
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(words[1]))
			fmt.Fprintf(conn, "Content-Type: text/html\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, words[1])
			i++
		}
	}
}
