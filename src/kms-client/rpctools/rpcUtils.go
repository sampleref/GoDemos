package rpctools

import (
	"golang.org/x/net/websocket"
	"log"
	"encoding/json"
	"kms-client/common"
)

var origin = "http://localhost/"
var ws_url = "ws://localhost:8888/kurento"

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
		"id":      common.GetRandomNumberId(),
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

func SendRequest(ws *websocket.Conn, request map[string]interface{}) {
	j, _ := json.MarshalIndent(request, "", "    ")
	log.Println("json", string(j))
	websocket.JSON.Send(ws, request)
}

func handleResponse(ws *websocket.Conn) {
	for { // run forever
		r := common.Response{}
		websocket.JSON.Receive(ws, &r)
		if r.Id != 0 {
			log.Println("handleResponse-Id: ", r.Id)
			if r.Result != nil {
				log.Println("handleResponse-Result: ", r.Result)
			}
			if r.Method != "" {
				log.Println("handleResponse-Method: ", r.Method)
			}
			if r.Params != nil {
				log.Println("handleResponse-Params: ",r.Params)
			}
			if r.Error != nil {
				log.Println("handleResponse-Error: ",r.Error)
			}
		}else {
			log.Println("handleResponse- No Id response ")
			if r.Error != nil {
				log.Println("handleResponse-Error: ",r.Error)
			}
		}
		if common.ResponseChannels[r.Id] != nil {
			common.ResponseChannels[r.Id] <- r
		} else {
			log.Println("Dropped message because there is no client ", r.Id)
			log.Println(r)
		}
	}
}
