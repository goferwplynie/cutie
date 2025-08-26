package cmd

import "github.com/spf13/cobra"

var (
	deadline string
	reminder int
)

var initCmd = &cobra.Command{
	Use:   "init [path] [name]",
	Short: "create a new project",
	Args:  cobra.ExactArgs(2),
	Run:   startProject,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&deadline, "dl", "", "project deadline (YYYY-MM-DD)")
	initCmd.Flags().IntVar(&reminder, "reminder", 0, "reminder in days")
}

func startProject(cmd *cobra.Command, args []string) {

}
