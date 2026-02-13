package main

import (
	"log"
	"template"
)

func main() {
	server, err := template.InitializeServer()
	if err != nil {
		log.Fatal(err)
	}

	server.Run(":5000")
}
