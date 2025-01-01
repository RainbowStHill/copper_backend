package identity_client_test

import (
	"testing"

	"github.com/rainbowsthill/copper_backend/pb"
	client "github.com/rainbowsthill/copper_backend/service/id/client"
)

func TestGetID(t *testing.T) {
	conn, err := pb.DailWithCreds("127.0.0.1", 17890, &pb.SerivceCertification{
		AppID:  "010000",
		AppKey: "identity_service",
	})
	if err != nil {
		t.Fatalf("Failed to connect to identity service: %v", err)
	}
	defer conn.Close()

	id, err := client.GetIdentity(conn, pb.GeneratorType_SNOWFLAKE)
	if err != nil {
		t.Fatalf("Failed to get identity: %v", err)
	}

	t.Logf("Got ID: %s", id)
}

func BenchmarkGetID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := pb.DailWithCreds("127.0.0.1", 17890, &pb.SerivceCertification{
			AppID:  "010000",
			AppKey: "identity_service",
		})
		if err != nil {
			b.Fatalf("Failed to connect to identity service: %v", err)
		}
		defer conn.Close()

		_, err = client.GetIdentity(conn, pb.GeneratorType_SNOWFLAKE)
		if err != nil {
			b.Fatalf("Failed to get identity: %v", err)
		}
	}
}
