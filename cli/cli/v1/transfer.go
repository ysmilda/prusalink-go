package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Retrieves information about the current transfer.",

	RunE: func(_ *cobra.Command, _ []string) error {
		transfer, err := conn.Transfer()
		if err != nil {
			return err
		}
		if transfer == nil {
			println("No transfer in progress")
			return nil
		}
		return cli.Print(transfer)
	},
}

func init() {
	v1Cmd.AddCommand(transferCmd)
}
