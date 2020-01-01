package main

import (
	"log"

	"github.com/wpjunior/friday-home/api"
	"github.com/wpjunior/friday-home/tv/samsung"
)

func main() {
	apiInstance := api.New(
		samsung.New(),
	)
	err := apiInstance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
