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
	go func() {
		for {
			time.Sleep(7 * time.Second)
			client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})
			var size int32 = 12
			resp, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: size})
			if err != nil {
				fmt.Printf("oh no: %v\n", err)
				continue
			}
			fmt.Printf("Sent: %v, responded with: %v \n", size, resp.Inches)
		}
	}()

	startServer()
}

func startServer() {
	srv := &server.Server{} // implements Haberdasher interface

	handler := helloworld.NewHelloWorldServer(srv)
	fmt.Println("Starting Twirp ServiceA...")
	http.ListenAndServe(":8081", handler)
}
