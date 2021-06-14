package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/wung-s/hellotwirp/internal/server"
	"github.com/wung-s/hellotwirp/rpc/haberdasher"
	"github.com/wung-s/hellotwirp/rpc/helloworld"
)

func main() {
	startClient()
	srv := &server.Server{} // implements Haberdasher interface
	handler := haberdasher.NewHaberdasherServer(srv)
	fmt.Println("Starting Twirp ServiceB...")
	http.ListenAndServe(":8080", handler)
}

func startClient() {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			client := helloworld.NewHelloWorldProtobufClient("http://localhost:8081", &http.Client{})
			greeting, err := client.Hello(context.Background(), &helloworld.HelloReq{Subject: "Testing..."})
			if err != nil {
				fmt.Printf("oh no: %v \n", err)
				continue
			}
			fmt.Printf("Received response: %+v\n", greeting)
		}
	}()
}
