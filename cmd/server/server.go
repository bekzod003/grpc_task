package main

import (
	"grpc_task/pkg/proto"
	"grpc_task/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
}

func contServ() {
	cs := grpc.NewServer()
	srv := server.GRPCContactSrever{}

	proto.RegisterContacListServer(cs, srv)

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = cs.Serve(ls); err != nil {
		log.Fatal(err)
	}
}

func taskServer() {
	ts := grpc.NewServer()
	srv := server.GRPCContactSrever{}

	proto.RegisterContacListServer(ts, srv)

	ls, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	if err = ts.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
