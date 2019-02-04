package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	pb "Makdyff/dialog-test/proto"

	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var (
	c pb.GreeterClient
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewGreeterClient(conn)

	logs.Info("Run test...")

	go tmp1()
	go tmp2()
	go tmp3()

	var gracefulStop = make(chan os.Signal)
	<-gracefulStop

	logs.Warn("Exit")
}

func tmp1() {
	for {
		time.Sleep(5 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		if t, err := c.GetInfo(ctx, &pb.UrlNameRequest{Name: "google.com"}); err != nil {
			logs.Error("could not greet: %v", err)
		} else {
			tt, _ := json.Marshal(t)
			logs.Warn(string(tt))
		}

		cancel()
	}
}

func tmp2() {
	for {
		time.Sleep(20 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		if t, err := c.GetMaxMinPing(ctx, &pb.Empty{}); err != nil {
			logs.Error("could not greet: %v", err)
		} else {
			tt, _ := json.Marshal(t)
			logs.Warn(string(tt))
		}

		cancel()
	}
}

func tmp3() {
	for {
		time.Sleep(20 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		if t, err := c.RequestStat(ctx, &pb.Empty{}); err != nil {
			logs.Error("could not greet: %v", err)
		} else {
			tt, _ := json.Marshal(t)
			logs.Warn(string(tt))
		}

		cancel()
	}
}
