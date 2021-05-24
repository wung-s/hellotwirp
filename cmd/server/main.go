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
	startClient()
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	handler := haberdasher.NewHaberdasherServer(server)
	fmt.Println("Starting Twirp Server on server/main...")
	http.ListenAndServe(":8080", handler)
}

func startClient() {
	go func() {
		for {
			time.Sleep(7 * time.Second)
			client := helloworld.NewHelloWorldProtobufClient("http://localhost:8081", &http.Client{})
			greeting, err := client.Hello(context.Background(), &helloworld.HelloReq{Subject: "Testing..."})
			if err != nil {
				fmt.Printf("oh no: %v \n", err)
				continue
			}
			fmt.Printf("I have a nice new hat: %+v", greeting)
		}
	}()
}
