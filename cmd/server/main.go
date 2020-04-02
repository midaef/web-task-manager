package main

import (
	"log"
	"packages/internal/app/server"
)

func main() {
	err := server.NewServer()
	if err != nil {
		log.Fatal("Server error", err)
	}
}
