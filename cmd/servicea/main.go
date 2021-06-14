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
		var size int32 = 10
		for {
			time.Sleep(3 * time.Second)
			client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})
			resp, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: size})
			if err != nil {
				fmt.Printf("oh no: %v\n", err)
				continue
			}
			fmt.Printf("Sent: %v, responded with: %v \n", size, resp.Inches)
			size += resp.Inches
		}
	}()

	startServer()
}

func startServer() {
	srv := &server.Server{} // implements Haberdasher interface

	mux := http.NewServeMux()
	rpcHandler := helloworld.NewHelloWorldServer(srv)
	mux.Handle(rpcHandler.PathPrefix(), rpcHandler)
	mux.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})
	fmt.Println("Starting Twirp ServiceA...")
	http.ListenAndServe(":8081", mux)
}
