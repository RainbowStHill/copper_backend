package identity_server

import (
	"context"
	"fmt"
	"net"

	"github.com/rainbowsthill/copper_backend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type IdentityServer struct {
	pb.UnimplementedIdentityServer
}

func (is *IdentityServer) GetID(ctx context.Context, req *pb.IDReq) (*pb.IDResp, error) {
	generator := GetIDGenerator(pb.GeneratorType_name[int32(req.GetGenerator())])
	return &pb.IDResp{
		Id: generator.Generate().String(),
	}, nil
}

func Run(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("Listener for port %d failed to initialize: %v", port, err)
	}

	// TLS certification
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to load credential files: %v", err)
	}

	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterIdentityServer(server, &IdentityServer{})

	server.Serve(listener)
}
