package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
)

var (
	forceMove bool

	movePadCmd = NewMovePadCmd()
)

func init() {
	rootCmd.AddCommand(movePadCmd)
}

func NewMovePadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "move-pad [sourceID] [destinationID]",
		Short: "Moves a single Pad",
		Long:  "The command moves a single pad. If force is true and the destination pad exists, it will be overwritten.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				cmd.Print(cmd.UsageString())
				return
			}

			etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
			sourceID := args[0]
			destinationID := args[1]

			err := etherpad.MovePad(sourceID, destinationID, forceMove)
			if err != nil {
				log.WithError(err).WithFields(log.Fields{"sourceID": sourceID, "destinationID": destinationID}).Error("error while moving pad")
			} else {
				log.WithFields(log.Fields{"sourceID": sourceID, "destinationID": destinationID}).Info("pad successfully moved")
			}
		},
	}

	cmd.Flags().BoolVar(&forceMove, "force", false, "If set and the destination pad exists, it will be overwritten.")

	return cmd
}
