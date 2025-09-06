package cmd

import (
	"fmt"

	"github.com/goferwplynie/cutie/logger"
	"github.com/spf13/cobra"
)

var unfreezeCmd = &cobra.Command{
	Use:   "unfreeze",
	Short: "unfreeze project",
	Long:  `unfreeze project. Makes reminders show up again :33`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unfreeze called")
		if err := SetFreeze(args[0], false); err != nil {
			logger.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(unfreezeCmd)
}
