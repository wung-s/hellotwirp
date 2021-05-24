package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/wung-s/hellotwirp/internal/haberdasherserver"
	"github.com/wung-s/hellotwirp/rpc/haberdasher"
	"github.com/wung-s/hellotwirp/rpc/helloworld"
)

func main() {
	go func() {
		for {
			time.Sleep(7 * time.Second)
			client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

			hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
			if err != nil {
				fmt.Printf("oh no: %v\n", err)
				continue
			}
			fmt.Printf("I have a nice new hat: %+v\n", hat)
		}
	}()

	startServer()
}

func startServer() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface

	handler := helloworld.NewHelloWorldServer(server)
	fmt.Println("Starting Twirp Server on client/main...")
	http.ListenAndServe(":8081", handler)
}
