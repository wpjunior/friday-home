package main

import (
	"log"

	"github.com/wpjunior/friday-home/api"
	"github.com/wpjunior/friday-home/player/vlc"
	"github.com/wpjunior/friday-home/tv/samsung"
)

func main() {
	apiInstance := api.New(
		samsung.New(),
		vlc.New(),
	)
	err := apiInstance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
