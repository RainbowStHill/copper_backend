package pb

import (
	context "context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type SerivceCertification struct {
	AppID  string
	AppKey string
}

// NewServiceCertification creates a new RPC certification.
func NewServiceCertification(appID, appKey string) *SerivceCertification {
	return &SerivceCertification{
		AppID:  appID,
		AppKey: appKey,
	}
}

func (sc SerivceCertification) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  sc.AppID,
		"appkey": sc.AppKey,
	}, nil
}

// RequrieTransportSecurity always return true.
func (sc SerivceCertification) RequireTransportSecurity() bool {
	return true
}

// DailWithCreds gets authorization information from ./keys/* and ./config_files/*.yaml and connect to specified target.
func DailWithCreds(addr string, port int, cert *SerivceCertification) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}

	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "www.copper.eu.org")
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials for %s: %v", cert.AppKey, err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	opts = append(opts, grpc.WithPerRPCCredentials(cert))

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", addr, port), opts...)
	if err != nil {
		return nil, err
	}
	return conn, err
}
