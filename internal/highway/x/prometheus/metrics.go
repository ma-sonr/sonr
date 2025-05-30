package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/kataras/golog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sonr-io/sonr/pkg/host"
)

type HighwayTelemetry struct {
	endpoint string
}

var (
	logger       = golog.Default.Child("node/telemetry")
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "Sonr",
		Name:      "total_events_processed",
		Help:      "The total number of processed events on the highway",
	})

	objectsAdded = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "Sonr",
		Name:      "objects",
		Subsystem: "data_storage",
		Help:      "Counts the number of objects added to the highway",
	})

	blobsAdded = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "Sonr",
		Name:      "files",
		Subsystem: "file_storage",
		Help:      "Counts the number of objects added to the highway",
	})
)

// `New` is a function that takes a context and a host, and returns a HighwayTelemetry struct and an
// error
func New(ctx context.Context, hn host.SonrHost) (*HighwayTelemetry, error) {
	defer RegisterEvents()
	return &HighwayTelemetry{endpoint: "/metrics"}, nil
}

// It's incrementing the `opsProcessed` counter every 2 seconds.
func (ht *HighwayTelemetry) RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

// It's returning a handler that will be used to serve the metrics.
func (ht *HighwayTelemetry) GetMetricsHandler() http.Handler {
	defer ht.RecordMetrics()
	return promhttp.Handler()
}
