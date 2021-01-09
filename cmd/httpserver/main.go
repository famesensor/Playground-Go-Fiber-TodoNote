package main

import (
	"log"

	"github.com/famesensor/playground-go-fiber-todonotes/protocol"
)

func main() {
	if err := protocol.ServerHttp(); err != nil {
		log.Fatal(err)
	}
}
