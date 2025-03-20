package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// abertura de conexao tcp ip
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// caso haja a conexao do client, joga o processo em thread que vai gerencia-lo
		//liberando a conexao de outros clientes
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		if _, err := io.WriteString(c, "Response form server\n"); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
