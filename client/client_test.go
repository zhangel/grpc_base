package client

import (
	"context"
	grpc "google.golang.org/grpc"
	"grpc_base/api"
	"grpc_base/api/hello"
	"log"
	"testing"
)

func TestSetInfo(t *testing.T) {
	host := "10.49.172.35:8877"
	log.Printf("SetInfo start\n")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Printf("err=%+v", err)
	}
	client := api.NewHelloServiceClient(conn)
	ctx := context.Background()
	req := &hello.SetInfo_Request{Query: "request test args"}
	Response, err := client.SetInfo(ctx, req)
	if err != nil {
		log.Printf("request setinfo error=%+v", err)
	}
	log.Printf("error=%d | msg=%s", Response.Error, Response.Msg)
}
