package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(ds io.Writer, src io.Reader) {
	if _, err := io.Copy(ds, src); err != nil {
		log.Fatal(err)
	}
}
