package main

import (
	"context"

	"server/endpoint"
	myGrpc "server/grpc"
	"server/pb"

	"github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	diff grpc.Handler
}

func (s *grpcServer) Diff(ctx context.Context, r *pb.StringRequest) (*pb.StringResponse, error) {
	_, resp, err := s.diff.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.StringResponse), nil

}

func NewGRPCServer(ctx context.Context, endpoints endpoint.StringEndpoints, serverTracer grpc.ServerOption) pb.StringServiceServer {
	return &grpcServer{
		diff: grpc.NewServer(
			endpoints.StringEndpoint,
			myGrpc.DecodeGRPCStringRequest,
			myGrpc.EncodeGRPCStringResponse,
			serverTracer,
		),
	}
}
