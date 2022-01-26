package main

import (
	"flag"
	grpc "google.golang.org/grpc"
	"grpc_base/api"
	"grpc_base/server/services"
	"grpc_base/utils/ini"
	"log"
	"net"
)

var config_type = flag.String("-config.file.type",
	"ini", "Config file type. optionals: auto, yaml, json, toml, ini, hocon")
var config_path = flag.String("config.file.path", "", "config file path.")

func main() {
	flag.Parse()
	grpcServer := grpc.NewServer()
	api.RegisterHelloServiceServer(grpcServer, new(services.HelloService))
	defaultPort := ":8877"
	if *config_path != "" {
		err := ini.Load(*config_path)
		if err != nil {
			log.Printf("load config file error=%+v\n", err)
			return
		}
		defaultPort = ini.String("grpc_port")
	}
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Printf("err=%+v\n", err)
	}
	log.Printf("Serving Grpc on 0.0.0.0%s\n", defaultPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("run grpc server error=%+v\n", err)
	}

}
