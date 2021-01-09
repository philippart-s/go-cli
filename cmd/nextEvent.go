package cmd

import (
	"github.com/spf13/cobra"
)

// nextEventCmd represents the nextEvent command
var nextEventCmd = &cobra.Command{
	Use:   "nextEvent",
	Short: "Liste des 10 prochains évènements",
	Run: func(cmd *cobra.Command, args []string) {
		printNextEvents()
	},
}

func init() {
	rootCmd.AddCommand(nextEventCmd)
}
