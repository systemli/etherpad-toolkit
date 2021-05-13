package cmd

import (
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/metrics"
)

var (
	listenAddr string
	suffixes   string

	metricsCmd = NewMetricsCmd()
)

func init() {
	rootCmd.AddCommand(metricsCmd)
}

func NewMetricsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Serves Pad related metrics",
		Long:  "The Command serves the count of pads grouped by suffix in Prometheus format.",
		Run: func(cmd *cobra.Command, args []string) {
			etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
			prometheus.MustRegister(metrics.NewPadCollector(etherpad, strings.Split(suffixes, ",")))

			http.Handle("/metrics", promhttp.Handler())
			log.Fatal(http.ListenAndServe(listenAddr, nil))
		},
	}

	cmd.Flags().StringVar(&listenAddr, "listen.addr", ":9012", "Address on which to expose metrics.")
	cmd.Flags().StringVar(&suffixes, "suffixes", "keep,temp", "Suffixes to group the pads.")

	return cmd
}
