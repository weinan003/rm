package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/rm/comm"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) CreateNewResourceContext(ctx context.Context, empty *pb.Empty) (reply *pb.CreateResponse, err error) {
	log.Println("CreateNewResourceContext")
	reply = &pb.CreateResponse{}
	reply.ResourceId = 1

	return reply, nil
}

func (s *server) GetSelfNetInfo(ctx context.Context, empty *pb.Empty) (ret *pb.Segment, err error) {
	log.Println("Get Self Net Info")
	ret = &pb.Segment{
		Hostname: "localhost",
		Address:  "127.0.0.1",
		Port:     5432,
	}
	return ret, err
}

func (s *server) AcquireResourceFromRM(ctx context.Context, acquire *pb.AcquireResRequest) (reply *pb.AcquireResResponse, err error) {
	log.Println("AcquireResourceFromRM")

	reply = &pb.AcquireResResponse{}
	reply.SegCount = 1
	reply.Vcore = 0.0625
	reply.MemoryMb = 256

	segment := &pb.Segment{}
	segment.ID = int32(0)
	segment.HdfsHostname = "localhost"
	segment.Port = 40000
	segment.Hostname = "localhost"
	segment.Address = "127.0.0.1"
	reply.Segments = append(reply.Segments, segment)

	err = nil
	return reply, err
}

func (s *server) ReleaseResourceFromRM(ctx context.Context, in *pb.ReleaseResRequest) (*pb.Empty, error) {

	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRMServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
