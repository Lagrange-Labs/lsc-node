package network

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	networkv2types "github.com/Lagrange-Labs/lagrange-node/network/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
)

const (
	// DefaultHealthCheckTimeout is the default timeout for the health check.
	DefaultHealthCheckTimeout = 25 // seconds
	// DefaultHealthCheckInterval is the default interval for the health check.
	DefaultHealthCheckInterval = 5 // seconds
	// DefaultHealthCheckRetry is the default retry for the health check.
	DefaultHealthCheckRetry = 3 // times
)

var (
	// ErrNoServerAvailable is the error when no server is available.
	ErrNoServerAvailable = errors.New("no server available")
	// ErrCurrentServerNotServing is the error when the current server is not serving.
	ErrCurrentServerNotServing = errors.New("the current server is not serving")
)

// healthManager is the struct to check the health of the server and switch the server.
type healthManager struct {
	serverURLs []string

	index   int
	conn    *grpc.ClientConn
	watcher grpc_health_v1.Health_WatchClient

	ctx    context.Context
	cancel context.CancelFunc
	chErr  chan error
}

// newHealthManager creates a new health manager.
func newHealthManager(serverURLs []string) (*healthManager, error) {
	ctx, cancel := context.WithCancel(context.Background())
	hm := &healthManager{
		serverURLs: serverURLs,
		index:      -1,
		ctx:        ctx,
		cancel:     cancel,
	}

	return hm, nil
}

// getHealthClient gets the health client.
func (hm *healthManager) getHealthClient() (networkv2types.NetworkServiceClient, error) {
	if err := hm.loadHealthClient(); err != nil {
		return nil, err
	}

	go hm.healthCheck()

	return networkv2types.NewNetworkServiceClient(hm.conn), nil
}

// loadHealthClient loads the health client.
func (hm *healthManager) loadHealthClient() error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	for index, serverURL := range hm.serverURLs {
		if index == hm.index {
			continue
		}

		conn, err := grpc.NewClient(serverURL, opts...)
		if err != nil {
			logger.Warnf("failed to connect to the server %s: %v", serverURL, err)
			continue
		}
		if loaded := func() bool {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*DefaultHealthCheckTimeout)
			defer cancel()
			watcher, err := grpc_health_v1.NewHealthClient(conn).Watch(ctx, &grpc_health_v1.HealthCheckRequest{})
			if err != nil {
				logger.Warnf("failed to watch the health of the server %s: %v", serverURL, err)
			}
			for i := 0; i < DefaultHealthCheckRetry; i++ {
				resp, err := watcher.Recv()
				if err != nil {
					logger.Warnf("failed to receive the health of the server %s: %v", serverURL, err)
					continue
				}
				if resp.Status == grpc_health_v1.HealthCheckResponse_SERVING {
					hm.index = index
					hm.conn = conn
					hm.watcher, err = grpc_health_v1.NewHealthClient(conn).Watch(hm.ctx, &grpc_health_v1.HealthCheckRequest{})
					if err != nil {
						logger.Warnf("failed to watch the health of the server %s: %v", serverURL, err)
						continue
					}
					return true
				}
				time.Sleep(time.Second * DefaultHealthCheckInterval)
			}

			return false
		}(); loaded {
			return nil
		}
	}

	return ErrNoServerAvailable
}

// healthCheck checks the health of the server.
func (hm *healthManager) healthCheck() {
	ticker := time.NewTicker(time.Second * DefaultHealthCheckInterval)
	defer ticker.Stop()

	retry := 0

	for {
		select {
		case <-hm.ctx.Done():
			hm.chErr <- hm.ctx.Err()
			return
		case <-ticker.C:
			for i := 0; i < DefaultHealthCheckRetry; i++ {
				ti := time.Now()
				resp, err := hm.watcher.Recv()
				if err == nil && resp.Status == grpc_health_v1.HealthCheckResponse_SERVING {
					telemetry.MeasureSince(ti, "network", "health_check")
					telemetry.SetGauge(1, "network", "current_health_server_index")
					continue
				}
				logger.Warnf("the server is not serving: %v", err)
				retry++
				if retry >= DefaultHealthCheckRetry {
					hm.chErr <- ErrCurrentServerNotServing
					return
				}
			}
		}
	}
}
