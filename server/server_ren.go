package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080") // or localhost:8080
	if err != nil {
		log.Fatal(err)
	}

	// To avoid code crash in later part,
	// so call "close()" and "defer" before the part with a high chance of failure.
	defer listener.Close()
	fmt.Println("Waiting for connection")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		// Send message to client
		fmt.Fprintln(conn, "Connection succeed.")

		// for loop + "go" = implement multi-connection
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from: " + remoteAddr)

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Println(remoteAddr, "disconnected")
			} else {
				log.Print(err.Error())
			}

			return // If error occur, close the connection, ex: client disconnects.
		}

		fmt.Printf("%v: %v\n", remoteAddr, msg)
		fmt.Fprintln(conn, "Host received your msg")
		conn.Write([]byte("Another way to send msg to client\n"))
	}
	conn.Close()
}
