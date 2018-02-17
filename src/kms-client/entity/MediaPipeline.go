package entity

import (
	"golang.org/x/net/websocket"
	"kms-client/rpctools"
	"kms-client/common"
	"log"
)

type MediaPipeline struct {
	Id              float64
	Value           string
	SessionId       string
	responseChannel chan common.Response
}

func CreateMediaPipeline(ws *websocket.Conn) (mediaPipeline *MediaPipeline) {
	mediaPipeline = &MediaPipeline{}
	mediaPipeline.Id = common.GetRandomNumberId()
	mediaPipeline.responseChannel = make(chan common.Response)
	common.ResponseChannels[mediaPipeline.Id] = mediaPipeline.responseChannel
	mediaPipeline.create(ws)
	go handleResponse(mediaPipeline)
	return
}

func (mediaPipeline *MediaPipeline) create(ws *websocket.Conn) {
	request := map[string]interface{}{
		"id":      mediaPipeline.Id,
		"jsonrpc": "2.0",
		"method":  "create",
		"params": map[string]interface{}{
			"type":              "MediaPipeline",
			"constructorParams": map[string]interface{}{},
			"properties":        map[string]interface{}{},
		},
	}
	rpctools.SendRequest(ws, request)
}

func handleResponse(mediaPipeline *MediaPipeline) {
	for {
		var response = <-mediaPipeline.responseChannel
		if response.Result != nil {
			log.Println("Received response at MediaPipeline: ", response.Result)
			mediaPipeline.SessionId = response.Result["sessionId"]
			mediaPipeline.Value = response.Result["value"]
		}
	}
}
