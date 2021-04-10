package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/systemli/etherpad-toolkit/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.WithError(err).Fatal("cannot start etherpad-toolkit")
	}
}
