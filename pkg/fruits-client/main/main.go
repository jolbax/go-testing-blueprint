package main

import (
	fruits_client "go-testing-blueprint/pkg/fruits-client"
	"log"
)

func main() {
	client := fruits_client.NewFruitsClient()
	someFruits, err := client.Get()
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}
	log.Println(someFruits.Fruits, someFruits.Code)

	otherFruits, err := client.Post()
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}
	log.Println(otherFruits.Fruits, otherFruits.Code)
}
