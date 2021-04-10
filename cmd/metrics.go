package cmd

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolchain/pkg"
	"github.com/systemli/etherpad-toolchain/pkg/metrics"
)

var (
	listenAddr string

	metricsCmd = &cobra.Command{
		Use:   "metrics",
		Short: "Serves Pad related metrics",
		Long:  "Command to serve pad related metrics for Prometheus",
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
