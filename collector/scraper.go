package collector

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

// Scraper is minimal interface that let's you add new prometheus metrics to prometheus-exporter.
type Scraper interface {
	// Name of the Scraper. Should be unique.
	Name() string

	// Help describes the role of the Scraper.
	Help() string

	// Version of xxx from which scraper is available.
	Version() float64

	// Scrape collects data from database connection and sends it over channel as prometheus metric.
	Scrape(ctx context.Context, ch chan<- prometheus.Metric,logger log.Logger) error
}
