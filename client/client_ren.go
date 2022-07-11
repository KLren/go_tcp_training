package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Print(err.Error())
	}
	reader := bufio.NewReader(conn)
	receiveMsg(reader)
	defer conn.Close()

	// Send msg
	conn.Write([]byte("Hello host\n"))
	receiveMsg(reader)
	fmt.Fprintf(conn, "Hello host by fmt\n")
	receiveMsg(reader)

	conn.Close()
}

func receiveMsg(r *bufio.Reader) {
	msg, err := r.ReadString('\n')
	if err != nil {
		log.Print(err.Error())
	} else {
		fmt.Println(msg)
	}
}
