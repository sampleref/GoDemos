package main

import (
	"fmt"
	"kms-client/rpctools"
	"log"
	"net/http"
	"os"
	"kms-client/entity"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println("Hello")
	ws, err := rpctools.CreateWebSocketClient()
	if err != nil {
		log.Fatal(err)
	}
	rpctools.PingWebSocket(ws)
	mediaPipeline := entity.CreateMediaPipeline(ws)
	time.Sleep(100 * time.Millisecond)
	log.Println("Mediapipeline session id: ", mediaPipeline.SessionId)
	http.ListenAndServe("localhost:7000", nil)
}

/*import "kms-client/test"
func main()  {
	test.TestRandomNumber()
}*/
