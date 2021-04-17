package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/helper"
	"github.com/systemli/etherpad-toolkit/pkg/purge"
)

var (
	concurrency int
	dryRun      bool
	expiration  string

	longDescription = `
The command checks every Pad for it’s last edited date. If it is older than the defined limit, the pad will be deleted.

Pads without any changes (revisions) will be deleted. This can happen when no content was changed in the pad 
(e.g. a person misspelles a pad).
Pads will grouped by the pre-defined suffixes. Every suffix has a defined expiration time. If the pad is older than the 
defined expiration time, the pad will be deleted.

Example:

etherpad-toolkit purge --expiration "default:720h,temp:24h,keep:8760h"

This configuration will group the pads in three clusters: default (expiration: 30 days, suffix is required!), 
temp (expiration: 24 hours), keep (expiration: 365 days). If pads in the clusters older than the given expiration the 
pads will be deleted.
`

	purgeCmd = &cobra.Command{
		Use:   "purge",
		Short: "Removes old Pads entirely from Etherpad",
		Long:  longDescription,
		Run:   runPurger,
	}
)

func init() {
	purgeCmd.Flags().StringVar(&expiration, "expiration", "", "Configuration for pad expiration duration. Example: \"default:720h,temp:24h,keep:8760h\"")
	purgeCmd.Flags().IntVar(&concurrency, "concurrency", 4, "Concurrency for the purge process")
	purgeCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Enable dry-run")

	rootCmd.AddCommand(purgeCmd)
}

func runPurger(cmd *cobra.Command, args []string) {
	etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
	exp, err := helper.ParsePadExpiration(expiration)
	if err != nil {
		log.WithError(err).Error("failed to parse expiration string")
		return
	}
	purger := purge.NewPurger(etherpad, exp, dryRun)
	purger.PurgePads(concurrency)
}
