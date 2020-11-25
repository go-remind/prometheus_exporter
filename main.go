package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"runtime"

	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"

	"github.com/go-remind/prometheus_exporter/collector"
)

var (
	// BuildVersion, BuildDate, Builder are filled in by the build script
	BuildVersion = "<<< filled in by build >>>"
	BuildDate    = "<<< filled in by build >>>"
)

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":8080", "The address to listen on for HTTP requests.")
		metricsPath   = flag.String("web.metircs-path", "/metrics", "Path under which to expose metrics.")
	)
	flag.Parse()
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	var scrapers = []collector.Scraper{
		collector.TypeDemo{},
	}

	buildInfo := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prometheus_exporter_build_info",
		Help: "prometheus exporter build_info",
	}, []string{"version", "build_date", "golang_version"})
	buildInfo.WithLabelValues(BuildVersion, BuildDate, runtime.Version()).Set(1)

	ctx := context.Background()
	prometheus.MustRegister(buildInfo)
	prometheus.MustRegister(collector.New(ctx, collector.NewMetrics(), scrapers, logger))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
<html>
<head><title>Redis Exporter v` + BuildVersion + `</title></head>
<body>
<h1>Redis Exporter ` + BuildVersion + `</h1>
<p><a href='` + *metricsPath + `'>Metrics</a></p>
</body>
</html>
`))
	})

	level.Info(logger).Log("msg", "Listening on address", "address", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
