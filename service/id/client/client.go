package identity_client

import (
	"context"
	"time"

	"github.com/rainbowsthill/copper_backend/pb"
	"google.golang.org/grpc"
)

func GetIdentity(conn *grpc.ClientConn, generatorType pb.GeneratorType) (string, error) {
	client := pb.NewIdentityClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetID(ctx, &pb.IDReq{
		Generator: generatorType,
	})
	if err != nil {
		return "", err
	}

	return resp.GetId(), nil
}
