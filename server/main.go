package main

import (
	"Makdyff/dialog-test/server/services"
	"net"

	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"

	pb "Makdyff/dialog-test/proto"
)

const (
	port = ":50051"
)

var (
	statS services.StatisticService
)

// server is used to implement helloworld.GreeterServer.
func main() {
	pingS := services.NewPingService()
	fileS := services.NewFileService("conf/sites.txt")
	statS := services.NewStatisticService()

	pingS.AddHandler(statS.Ping)
	fileS.AddHandler(pingS.UrlString)

	pingS.Run()
	fileS.Run()
	statS.Run()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logs.Error("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		logs.Error("failed to serve: %v", err)
	}

	logs.Info("Run test...")
}
