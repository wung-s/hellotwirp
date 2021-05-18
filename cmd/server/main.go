package main

import (
	"fmt"
	"net/http"

	"github.com/wung-s/hellotwirp/internal/haberdasherserver"
	"github.com/wung-s/hellotwirp/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)
	fmt.Println("Starting Twirp Server...")
	http.ListenAndServe(":8080", twirpHandler)
}
