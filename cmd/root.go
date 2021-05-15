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

	cmd.PersistentFlags().String("etherpad.url", "http://localhost:9001", "URL to access Etherpad (Env: ETHERPAD_URL)")
	cmd.PersistentFlags().String("etherpad.apikey", "", "API Key for Etherpad (Env: ETHERPAD_APIKEY)")
	cmd.PersistentFlags().String("log.level", "info", "Log level (Env: LOG_LEVEL)")
	cmd.PersistentFlags().String("log.format", "text", "Format for log output (Env: LOG_FORMAT)")

	etherpadUrl = os.Getenv("ETHERPAD_URL")
	if etherpadUrl == "" {
		etherpadUrl = cmd.Flag("etherpad.url").Value.String()
	}

	etherpadApiKey = os.Getenv("ETHERPAD_APIKEY")
	if etherpadApiKey == "" {
		etherpadApiKey = cmd.Flag("etherpad.apikey").Value.String()
	}

	logLevel = os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = cmd.Flag("log.level").Value.String()
	}

	logFormat = os.Getenv("LOG_FORMAT")
	if logFormat == "" {
		logFormat = cmd.Flag("log.format").Value.String()
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
