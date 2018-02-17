package main

import (
	"fmt"
	"kms-client/rpctools"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println("Hello")
	ws, err := rpctools.CreateWebSocketClient()
	if err != nil {
		log.Fatal(err)
	}
	rpctools.PingWebSocket(ws)
	http.ListenAndServe("localhost:7000", nil)
}
