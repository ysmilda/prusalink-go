package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Retrieves storage information about the printer.",
	RunE: func(_ *cobra.Command, _ []string) error {
		storage, err := conn.Storage()
		if err != nil {
			return err
		}
		return cli.Print(storage)
	},
}

func init() {
	v1Cmd.AddCommand(storageCmd)
}
