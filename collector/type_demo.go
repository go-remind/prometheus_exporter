package collector

import (
	"context"
	"math/rand"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

// Metric descriptors.
var (
	counterTypeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "type", "counter"),
		"The demo of counter.",
		[]string{}, nil,
	)
	gaugeTypeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "type", "gauge"),
		"The demo of gauge.",
		[]string{}, nil,
	)
	histogramTypeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "type", "histogram"),
		"The demo of histogram.",
		[]string{"status", "method"}, nil,
	)
	summaryTypeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "type", "summary"),
		"The demo of summary.",
		[]string{"status", "method"}, nil,
	)
)

type TypeDemo struct{}

// Name of the Scraper. Should be unique.
func (TypeDemo) Name() string {
	return "type"
}

// Help describes the role of the Scraper.
func (TypeDemo) Help() string {
	return "Collect the type"
}

// Version of xxx from which scraper is available.
func (TypeDemo) Version() float64 {
	return 2.0
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (TypeDemo) Scrape(ctx context.Context, ch chan<- prometheus.Metric, logger log.Logger) error {
	// Counter
	ch <- prometheus.MustNewConstMetric(
		counterTypeDesc, prometheus.CounterValue, float64(1),
	)
	// Gauge
	ch <- prometheus.MustNewConstMetric(
		gaugeTypeDesc, prometheus.GaugeValue, float64(rand.Intn(10)),
	)
	// Histogram
	ch <- prometheus.MustNewConstHistogram(
		histogramTypeDesc,
		4771, 403.34,
		map[float64]uint64{0.5: 42, 0.9: 323},
		"200", "get",
	)
	// Summary
	ch <- prometheus.MustNewConstSummary(
		summaryTypeDesc,
		4771, 403.34,
		map[float64]float64{25: 121, 50: 2403, 70: 3221, 90: 4323},
		"200", "get",
	)
	return nil
}

// check interface
var _ Scraper = TypeDemo{}
