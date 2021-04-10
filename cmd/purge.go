package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolchain/pkg"
	"github.com/systemli/etherpad-toolchain/pkg/purge"
)

var (
	concurrency int
	dryRun      bool

	purgeCmd = &cobra.Command{
		Use:   "purge",
		Short: "Removes old Pads entirely from Etherpad",
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
	sorted := pkg.SortPads(pads)

	purger.PurgePads(sorted, concurrency)
}
