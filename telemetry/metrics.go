package telemetry

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-metrics"
	metricsprom "github.com/hashicorp/go-metrics/prometheus"

	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// globalLabels defines the set of global labels that will be applied to all
// metrics emitted using the telemetry package function wrappers.
var globalLabels = []metrics.Label{}

// Metrics supported format types.
const (
	FormatDefault    = ""
	FormatPrometheus = "prometheus"
	FormatText       = "text"
)

// Config defines the configuration options for application telemetry.
type Config struct {
	// Prefixed with keys to separate services
	ServiceName string `mapstructure:"ServiceName"`

	// PrometheusRetentionTime, when positive, enables a Prometheus metrics sink.
	// It defines the retention duration in seconds. If 0, Prometheus metrics are
	// disabled.
	PrometheusRetentionTime utils.TimeDuration `mapstructure:"PrometheusRetentionTime"`
}

// NewGlobal creates a global Metrics.
func NewGlobal(cfg Config) (err error) {
	if cfg.PrometheusRetentionTime <= 0 {
		return fmt.Errorf("metrics is not enabled")
	}

	metricsConf := metrics.DefaultConfig(cfg.ServiceName)
	metricsConf.EnableHostnameLabel = true
	metricsConf.EnableServiceLabel = true
	prometheusOpts := metricsprom.PrometheusOpts{
		Expiration: time.Duration(cfg.PrometheusRetentionTime),
	}
	promSink, err := metricsprom.NewPrometheusSinkFrom(prometheusOpts)
	if err != nil {
		return err
	}
	if _, err := metrics.NewGlobal(metricsConf, promSink); err != nil {
		return err
	}

	return nil
}

// SetLabels sets the global labels.
func SetLabel(label metrics.Label) {
	globalLabels = append(globalLabels, label)
}
