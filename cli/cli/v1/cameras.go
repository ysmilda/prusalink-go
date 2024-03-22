package v1

import (
	"github.com/spf13/cobra"
)

var camerasCmd = &cobra.Command{
	Use:   "cameras",
	Short: "Retrieves information about the cameras on the printer.",
	Run: func(cmd *cobra.Command, _ []string) {
		println("Not implemented yet due to lack of hardware for testing.")
	},
}

func init() {
	v1Cmd.AddCommand(camerasCmd)
}
