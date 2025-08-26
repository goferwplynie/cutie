package cmd

import (
	"fmt"
	"os"

	"github.com/goferwplynie/cutie/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cutie",
	Short: "cute project manager :3",
	Long:  "cli tool to manage your whole projects (creating with templates, managing deadlines, reminders)",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&logger.Verbose, "verbose", "v", false, "enable verbose mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
