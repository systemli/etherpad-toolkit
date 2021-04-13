package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
)

var (
	deletePadCmd = &cobra.Command{
		Use:   "delete-pad [pad]",
		Short: "Removes a single Pad",
		Long:  "The command removes a single pad entirely from Etherpad.",
		Run:   runDelete,
	}
)

func init() {
	rootCmd.AddCommand(deletePadCmd)
}

func runDelete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(cmd.UsageString())
		return
	}

	etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
	pad := args[0]

	err := etherpad.DeletePad(pad)
	if err != nil {
		log.WithError(err).WithField("pad", pad).Error("error while deleting pad")
	} else {
		log.WithField("pad", pad).Info("pad successfully deleted")
	}
}
