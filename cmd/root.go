package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	etherpadUrl    string
	etherpadApiKey string
	logLevel       string
	logFormat      string

	rootCmd = &cobra.Command{
		Use:   "etherpad-toolkit",
		Short: "A toolkit for Etherpad",
		Long:  "Etherpad Toolchain is a collection for most common Etherpad maintenance tasks.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&etherpadUrl, "etherpad.url", "http://localhost:9001", "URL to access Etherpad")
	rootCmd.PersistentFlags().StringVar(&etherpadApiKey, "etherpad.apikey", "", "API Key for Etherpad")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log.level", "info", "Log level")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log.format", "text", "Format for log output")

	if logFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).Fatal("failed to parse log level")
	}
	log.SetLevel(lvl)
}
