package main

import (
	pb "Makdyff/dialog-test/proto"
	"context"
)

type Server struct {
}

func (s *Server) GetInfo(c context.Context, in *pb.UrlNameRequest) (*pb.UrlNameReply, error) {
	t := statS.GetInfo(in.Name)

	return &pb.UrlNameReply{
		Url:        t.Url,
		StatusCode: int32(t.StatusCode),
		PingTime:   t.PingTime.Nanoseconds(),
	}, nil
}

func (s *Server) GetMaxMinPing(context.Context, *pb.Empty) (*pb.MaxMinPingReply, error) {
	t := statS.GetMaxMinPing()

	return &pb.MaxMinPingReply{
		Max: &pb.MaxMinPingReply_PhoneNumber{
			Url:      t.Max.Url,
			PingTime: t.Max.PingTime.Nanoseconds(),
		},
		Min: &pb.MaxMinPingReply_PhoneNumber{
			Url:      t.Min.Url,
			PingTime: t.Min.PingTime.Nanoseconds(),
		},
	}, nil
}

func (s *Server) RequestStat(context.Context, *pb.Empty) (*pb.RequestStatReply, error) {
	t := statS.RequestStat()
	m := &pb.RequestStatReply{
		RequestStat1: make([]*pb.UrlNameReply, 0),
		RequestStat2: make([]*pb.MaxMinPingReply, 0),
	}

	for _, v1 := range t.RequestStat1 {
		m.RequestStat1 = append(m.RequestStat1, &pb.UrlNameReply{
			Url:        v1.Url,
			StatusCode: int32(v1.StatusCode),
			PingTime:   v1.PingTime.Nanoseconds(),
		})
	}

	for _, v2 := range t.RequestStat2 {
		m.RequestStat2 = append(m.RequestStat2, &pb.MaxMinPingReply{
			Max: &pb.MaxMinPingReply_PhoneNumber{
				Url:      v2.Max.Url,
				PingTime: v2.Max.PingTime.Nanoseconds(),
			},
			Min: &pb.MaxMinPingReply_PhoneNumber{
				Url:      v2.Min.Url,
				PingTime: v2.Min.PingTime.Nanoseconds(),
			},
		})
	}

	return m, nil
}
