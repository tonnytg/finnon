package main

import (
	"finnon/internal/infra/webserv"
	"log"
)

func main() {
	log.Println("Start Finnon")

	webserv.Start()

}
