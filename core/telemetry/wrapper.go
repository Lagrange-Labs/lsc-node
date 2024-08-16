package telemetry

import (
	"time"

	"github.com/hashicorp/go-metrics"
)

// NewLabel creates a new instance of Label with name and value
func NewLabel(name, value string) metrics.Label {
	return metrics.Label{Name: name, Value: value}
}

// IncrCounter provides a wrapper functionality for emitting a counter metric with
// global labels (if any).
func IncrCounter(val float32, keys ...string) {
	metrics.IncrCounterWithLabels(keys, val, globalLabels)
}

// IncrCounterWithLabels provides a wrapper functionality for emitting a counter
// metric with global labels (if any) along with the provided labels.
func IncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.IncrCounterWithLabels(keys, val, append(labels, globalLabels...))
}

// SetGauge provides a wrapper functionality for emitting a gauge metric with
// global labels (if any).
func SetGauge(val float64, keys ...string) {
	metrics.SetPrecisionGaugeWithLabels(keys, val, globalLabels)
}

// SetGaugeWithLabels provides a wrapper functionality for emitting a gauge
// metric with global labels (if any) along with the provided labels.
func SetGaugeWithLabels(keys []string, val float64, labels []metrics.Label) {
	metrics.SetPrecisionGaugeWithLabels(keys, val, append(labels, globalLabels...))
}

// AddSample provides a wrapper functionality for emitting a sample metric with
// global labels (if any).
func AddSample(val float32, keys ...string) {
	metrics.AddSampleWithLabels(keys, val, globalLabels)
}

// AddSampleWithLabels provides a wrapper functionality for emitting a sample
// metric with global labels (if any) along with the provided labels.
func AddSampleWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.AddSampleWithLabels(keys, val, append(labels, globalLabels...))
}

// MeasureSince provides a wrapper functionality for emitting a a time measure
// metric with global labels (if any).
func MeasureSince(start time.Time, keys ...string) {
	metrics.MeasureSinceWithLabels(keys, start.UTC(), globalLabels)
}
