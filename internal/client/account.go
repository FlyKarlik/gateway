package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gateway/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	pb "protos/account"
	"time"
)

// AccountClient client for account service
type AccountClient struct {
	config *config.Config
}

// NewAccountClient create new client
func NewAccountClient(config *config.Config) *AccountClient {
	return &AccountClient{
		config: config,
	}
}

func (cl *AccountClient) ConnectToAccountService(ctx context.Context) (context.Context, context.CancelFunc, pb.AccountServiceClient, *grpc.ClientConn, error) {
	b, err := os.ReadFile("cert/ca.cert")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed read file cert: %w", err)
	}
	cp := x509.NewCertPool()

	if !cp.AppendCertsFromPEM(b) {
		log.Println("credentials: failed to append certificates")
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}

	conn, err := grpc.Dial(
		cl.config.AccountServiceHost,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("grpc.Dial error: %w", err)
	}
	c := pb.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	return ctx, cancel, c, conn, nil
}
