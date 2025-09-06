package cmd

import (
	"github.com/goferwplynie/cutie/logger"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:   "freeze",
	Short: "freeze project",
	Long: `freeze project to not show reminders for it.
	Use it when you know that you won't be working on it for some time and don't want to see reminders.
	You can unfreeze it at any time
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := SetFreeze(args[0], true); err != nil {
			logger.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(freezeCmd)
}
