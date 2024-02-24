package main

import (
	"log"

	"github.com/PetarPeychev/test-signer-service/api"
)

func main() {
	config, err := api.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
