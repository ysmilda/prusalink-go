package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieves status information about the printer.",
	RunE: func(_ *cobra.Command, _ []string) error {
		status, err := conn.Status()
		if err != nil {
			return err
		}
		return cli.Print(status)
	},
}

func init() {
	v1Cmd.AddCommand(statusCmd)
}
