package HelloServiceRPC

import (
	"HelloDemo/proto"
	"golang.org/x/net/context"
	"fmt"
	"bytes"
)

type HelloServiceServer struct {
	NamePrefix string
}

func (helloService HelloServiceServer) GetHelloMessage(ctx context.Context, request *service_Hello.ServiceHelloRequest) (*service_Hello.ServiceHelloResponse, error) {
	fmt.Println(" Saying hello to " + request.SayHelloTo + " with prefix " + helloService.NamePrefix)
	var message bytes.Buffer
	message.WriteString("Hello ")
	message.WriteString(helloService.NamePrefix)
	message.WriteString(" ")
	message.WriteString(request.SayHelloTo)
	message.WriteString(" ! ")
	response := &service_Hello.ServiceHelloResponse{HelloMessage:message.String()}
	return response, nil
}