package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/systemli/etherpad-toolkit/pkg"
)

var (
	forceCopy bool

	copyPadCmd = NewCopyPadCmd()
)

func init() {
	rootCmd.AddCommand(copyPadCmd)
}

func NewCopyPadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy-pad [sourceID] [destinationID]",
		Short: "Copies a single Pad",
		Long:  "The command copies a pad with full history and chat. If force is true and the destination pad exists, it will be overwritten.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				cmd.Print(cmd.UsageString())
				return
			}

			etherpad := pkg.NewEtherpadClient(etherpadUrl, etherpadApiKey)
			sourceID := args[0]
			destinationID := args[1]

			err := etherpad.CopyPad(sourceID, destinationID, forceCopy)
			if err != nil {
				log.WithError(err).WithFields(log.Fields{"sourceID": sourceID, "destinationID": destinationID}).Error("error while copy pad")
			} else {
				log.WithFields(log.Fields{"sourceID": sourceID, "destinationID": destinationID}).Info("pad successfully copied")
			}
		},
	}

	cmd.Flags().BoolVar(&forceCopy, "force", false, "If set and the destination pad exists, it will be overwritten.")

	return cmd
}
