package main

import (
	"log"

	"github.com/williepotgieter/order-packs-calculator/internal/adapters/http"
)

func main() {
	server, err := http.NewAdapter()
	if err != nil {
		log.Fatalln("unable to setup server:", err.Error())
	}

	if err := server.Run(); err != nil {
		log.Fatalln("server experienced an error:", err.Error())
	}
}
