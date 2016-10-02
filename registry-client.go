package main

import (
	"fmt"

	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url := "localhost:5000"
	username := ""       // anonymous
	password := "" // anonymous
	hub, err := registry.NewInsecure(url, username, password)
	if err != nil {
		panic(err)
	}

	repositories, err := hub.Repositories()
	if err != nil {
		panic(err)
	}

	for _, repository := range repositories {
		tags, err := hub.Tags(repository)
		if err != nil {
			panic(err)
		}

		fmt.Println(repository)
		for _, tag := range tags {
			fmt.Println(" ", tag)
		}
	}
}
