package grpc

import (
	endpts "server/endpoint"
	"server/service"

	"server/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func StringDiff(conn *grpc.ClientConn, clientTracer kitgrpc.ClientOption) service.Service {

	var ep = grpctransport.NewClient(conn,
		"pb.StringService",
		"Diff",
		EncodeGRPCStringRequest,
		DecodeGRPCStringResponse,
		pb.StringResponse{},
		clientTracer,
	).Endpoint()

	StringEp := endpts.StringEndpoints{
		StringEndpoint: ep,
	}
	return StringEp
}
