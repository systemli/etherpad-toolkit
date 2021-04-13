package cmd

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/metrics"
)

var (
	listenAddr string

	metricsCmd = &cobra.Command{
		Use:   "metrics",
		Short: "Serves Pad related metrics",
		Long:  "The Command serves the count of pads grouped by suffix in Prometheus format.",
		Run:   runMetrics,
	}
)

func init() {
	metricsCmd.LocalFlags().StringVar(&listenAddr, "listen.addr", ":9012", "")

	rootCmd.AddCommand(metricsCmd)
}

func runMetrics(cmd *cobra.Command, args []string) {
	etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
	prometheus.MustRegister(metrics.NewPadCollector(etherpad))

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
