package main

import (
	"finnon/internal/webserv"
	"log"
)

func main() {
	log.Println("Start Finnon")

	webserv.Start()

}
