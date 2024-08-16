package signer

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/signer/types"
)

const (
	// PublicKeyMethod is the method to get the public key.
	PublicKeyMethod = "public_key"
)

var (
	// ErrSignerNotFound is returned when the signer is not found.
	ErrSignerNotFound = errors.New("signer not found")
)

// RunServer runs the signer server.
func RunServer(port string, certCfg *core.CertConfig, signers map[string]Signer) error {
	ctx := context.Background()

	signerService, err := NewSignerService(signers)
	if err != nil {
		return err
	}

	go func() {
		_ = runGRPCServer(ctx, signerService, port, certCfg)
	}()

	return nil
}

// HealthChecker will provide an implementation of the HealthCheck interface.
type healthChecker struct{}

// NewHealthChecker returns a health checker according to standard package
// grpc.health.v1.
func newHealthChecker() *healthChecker {
	return &healthChecker{}
}

// HealthCheck interface implementation.

// Check returns the current status of the server for unary gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch returns the current status of the server for stream gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func runGRPCServer(ctx context.Context, svc types.SignerServiceServer, port string, certCfg *core.CertConfig) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	if certCfg != nil {
		tlsConfig, err := core.LoadTLS(certCfg, true)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}
	server := grpc.NewServer(opts...)
	types.RegisterSignerServiceServer(server, svc)

	healthService := newHealthChecker()
	grpc_health_v1.RegisterHealthServer(server, healthService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	logger.Info("Signer Server is serving at ", port)
	return server.Serve(listen)
}

type signerService struct {
	types.UnimplementedSignerServiceServer
	signers map[string]Signer
}

// NewSignerService creates the signer service.
func NewSignerService(signers map[string]Signer) (types.SignerServiceServer, error) {
	return &signerService{
		signers: signers,
	}, nil
}

// Sign signs the message.
func (s *signerService) Sign(ctx context.Context, req *types.SignRequest) (*types.SignResponse, error) {
	logger.Infof("Sign request: %+v", req)
	signer, ok := s.signers[req.AccountId]
	if !ok {
		return nil, ErrSignerNotFound
	}

	if req.SignMethod == PublicKeyMethod {
		pubKey, err := signer.GetPubKey()
		if err != nil {
			return nil, err
		}

		return &types.SignResponse{
			Signature: common.Bytes2Hex(pubKey),
		}, nil
	}

	sig, err := signer.Sign(common.Hex2Bytes(req.Message))
	if err != nil {
		return nil, err
	}

	return &types.SignResponse{
		Signature: common.Bytes2Hex(sig),
	}, nil
}
