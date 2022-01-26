package services

import (
	"context"
	"grpc_base/api/hello"
	"log"
)

type HelloService struct{}

func (h *HelloService) SetInfo(ctx context.Context,
	req *hello.SetInfo_Request) (*hello.SetInfo_Response, error) {
	log.Printf("req=%+v", req)
	RetData := &hello.SetInfo_Response{Error: int32(1), Msg: "Fail"}
	return RetData, nil
}

func (h *HelloService) GetInfo(ctx context.Context, req *hello.GetInfo_Request) (*hello.GetInfo_Response, error) {
	RetData := &hello.GetInfo_Response{}
	return RetData, nil
}
