package rpctools

import (
	"golang.org/x/net/websocket"
	"log"
	"encoding/json"
	"kms-client/entity"
)

var origin string = "http://10.71.11.155/"
var ws_url string = "ws://10.71.11.155:8888/kurento"

func CreateWebSocketClient() (ws *websocket.Conn, err error) {
	ws, err = websocket.Dial(ws_url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	go handleResponse(ws)
	return
}

func PingWebSocket(ws *websocket.Conn) {
	request := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "ping",
		"params": map[string]interface{}{
			"interval": 240000,
		},
	}
	j, _ := json.MarshalIndent(request, "", "    ")
	log.Println("json", string(j))
	websocket.JSON.Send(ws, request)
}

func handleResponse(ws *websocket.Conn) {
	for { // run forever
		r := entity.Response{}
		websocket.JSON.Receive(ws, &r)
		if r.Result != nil {
			log.Println(r.Result)
		}
		if r.Params != nil {
			log.Println(r.Params)
		}
		if r.Error != nil {
			log.Println(r.Error)
		}
	}
}
