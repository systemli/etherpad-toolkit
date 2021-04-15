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

	longDescription = `
The command checks every Pad if the last edited date is older than the defined limit. Older Pads will be deleted.

Pads without any changes (revisions) will be deleted.
Pads without a suffix will be deleted after 30 days of inactivity.
Pads with the suffix "-temp" will be deleted after 24 hours of inactivity.
Pads with the suffix "-keep" will be deleted after 365 days of inactivity.
`

	purgeCmd = &cobra.Command{
		Use:   "purge",
		Short: "Removes old Pads entirely from Etherpad",
		Long:  longDescription,
		Run:   runPurger,
	}
)

func init() {
	purgeCmd.LocalFlags().IntVar(&concurrency, "concurrency", 4, "Concurrency for the purge process")
	purgeCmd.LocalFlags().BoolVar(&dryRun, "dry-run", false, "Enable dry-run")

	rootCmd.AddCommand(purgeCmd)
}

func runPurger(cmd *cobra.Command, args []string) {
	etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
	purger := purge.NewPurger(etherpad, dryRun)

	pads, err := etherpad.ListAllPads()
	if err != nil {
		log.WithError(err).Error("failed to fetch pads")
		return
	}
	sorted := helper.SortPads(pads)

	purger.PurgePads(sorted, concurrency)
}
