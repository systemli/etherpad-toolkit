package main

import (
	"flag"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	etherpadUrl    = flag.String("etherpad.url", "http://localhost:9001", "URL to access Etherpad")
	etherpadApiKey = flag.String("etherpad.apikey", "", "API Key for Etherpad")
	concurrency    = flag.Int("concurrency", 4, "Concurrency for the purge process")
	logFormat      = flag.String("log.format", "text", "Format for log output")
	logLevel       = flag.String("log.level", "info", "Log level")
)

func init() {
	flag.Parse()

	if *logFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.WithError(err).Fatal("failed to parse log level")
	}
	log.SetLevel(lvl)
}

func main() {
	etherpad := NewEtherpadClient(*etherpadUrl, *etherpadApiKey)
	purger := NewPurger(etherpad)

	pads, err := etherpad.ListAllPads()
	if err != nil {
		log.WithError(err).Error("failed to fetch pads")
		return
	}
	sorted := sortPads(pads)

	purger.PurgePads(sorted, *concurrency)
}

// sortPads will put the padIds into a string map organized by their suffixes
func sortPads(padIds []string) map[string][]string {
	sorted := make(map[string][]string)

	for _, pad := range padIds {
		if strings.HasSuffix(pad, "-keep") {
			sorted["keep"] = append(sorted["keep"], pad)
		} else if strings.HasSuffix(pad, "-temp") {
			sorted["temp"] = append(sorted["temp"], pad)
		} else {
			sorted["none"] = append(sorted["none"], pad)
		}
	}

	return sorted
}
