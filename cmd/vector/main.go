package main

import (
	"log"

	"github.com/jumuia/vector/internal/server"
)

func main() {
	cfg := server.Config{}
	err := server.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully exited")
}
