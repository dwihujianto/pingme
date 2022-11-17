package main

import (
	"log"

	"pingme/app"

)

func main() {
	ch := make(chan error)

	go func(c chan error) {
		c <- app.StartGrpc("localhost", 12345)
	}(ch)

	log.Println("Listen bro")

	for a := range ch {
		log.Println(a)
	}
}
