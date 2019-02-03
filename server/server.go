package main

import (
	pb "Makdyff/dialog-test/proto"
	"context"

	"github.com/astaxie/beego/logs"
)

type Server struct {
}

func (s *Server) GetInfo(c context.Context, in *pb.UrlNameRequest) (*pb.UrlNameReply, error) {
	t := statS.GetInfo(in.Name)
	logs.Info(in.Name, t)

	return &pb.UrlNameReply{
		Url:        t.Url,
		StatusCode: int32(t.StatusCode),
		PingTime:   t.PingTime.Nanoseconds(),
	}, nil
}

func (s *Server) GetMaxMinPing(context.Context, *pb.Empty) (*pb.MaxMinPingReply, error) {
	return &pb.MaxMinPingReply{Name: "Hello "}, nil
}

func (s *Server) RequestStat(context.Context, *pb.Empty) (*pb.RequestStatReply, error) {
	return &pb.RequestStatReply{Name: "Hello "}, nil
}
