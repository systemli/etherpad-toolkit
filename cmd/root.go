package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	etherpadUrl    string
	etherpadApiKey string
	logLevel       string
	logFormat      string

	rootCmd = NewRootCmd()
)

func Execute() error {
	return rootCmd.Execute()
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etherpad-toolkit",
		Short: "A toolkit for Etherpad",
		Long:  "Etherpad Toolkit is a collection for most common Etherpad maintenance tasks.",
	}

	cmd.PersistentFlags().StringVar(&etherpadUrl, "etherpad.url", "http://localhost:9001", "URL to access Etherpad (Env: ETHERPAD_URL)")
	cmd.PersistentFlags().StringVar(&etherpadApiKey, "etherpad.apikey", "", "API Key for Etherpad (Env: ETHERPAD_APIKEY)")
	cmd.PersistentFlags().StringVar(&logLevel, "log.level", "info", "Log level (Env: LOG_LEVEL)")
	cmd.PersistentFlags().StringVar(&logFormat, "log.format", "text", "Format for log output (Env: LOG_FORMAT)")

	if os.Getenv("ETHERPAD_URL") != "" {
		etherpadUrl = os.Getenv("ETHERPAD_URL")
	}

	if os.Getenv("ETHERPAD_APIKEY") != "" {
		etherpadApiKey = os.Getenv("ETHERPAD_APIKEY")
	}

	if os.Getenv("LOG_LEVEL") != "" {
		logLevel = os.Getenv("LOG_LEVEL")
	}

	if os.Getenv("LOG_FORMAT") != "" {
		logFormat = os.Getenv("LOG_FORMAT")
	}

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

	return cmd
}
