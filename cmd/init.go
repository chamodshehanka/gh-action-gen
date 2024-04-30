package cmd

import (
	internal "github.com/chamodshehanka/gh-action-gen/internal/init"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your workflows",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.LoadInitForm()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
