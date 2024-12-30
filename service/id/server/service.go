package identity_server

import (
	"context"
	"fmt"
	"net"
	"path/filepath"

	"github.com/rainbowsthill/copper_backend/config"
	"github.com/rainbowsthill/copper_backend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	opts := []grpc.ServerOption{}

	// TLS certification
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to load credential files: %v", err)
	}
	opts = append(opts, grpc.Creds(creds))

	// Authorizing information checker
	var authChecker = func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp any, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(authChecker))

	server := grpc.NewServer(opts...)

	pb.RegisterIdentityServer(server, &IdentityServer{})

	server.Serve(listener)
}

func auth(ctx context.Context) error {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Aborted, "Authorize infomation not found")
	}
	if appids, ok := meta["appid"]; !ok || len(appids) <= 0 {
		return status.Errorf(codes.Aborted, "Authorize infomation not found")
	}
	if appkeys, ok := meta["appkey"]; !ok || len(appkeys) <= 0 {
		return status.Errorf(codes.Aborted, "Authorize infomation not found")
	}

	var appid, appkey *string
	var err error = nil

	cf, _ := filepath.Abs("./config_files/config.yaml")

	if appid, err = config.GetBuiltInTypeConfig[string](cf, []string{"service", "appid"}); err != nil {
		return status.Errorf(codes.Aborted, "Server error")
	}
	if appkey, err = config.GetBuiltInTypeConfig[string](cf, []string{"service", "appkey"}); err != nil {
		return status.Errorf(codes.Aborted, "Server error")
	}

	if *appid != meta["appid"][0] || *appkey != meta["appkey"][0] {
		return status.Errorf(codes.Aborted, "Authorize infomation illegal")
	}

	return nil
}
