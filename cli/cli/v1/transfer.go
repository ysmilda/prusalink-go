package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Retrieves information about the current transfer.",

	RunE: func(_ *cobra.Command, _ []string) error {
		transfer, err := printer.Transfer().GetInfo()
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

	transferCmd.AddCommand(transferStopCmd)

	transferStopCmd.Flags().IntP("id", "i", 0, "ID of the transfer")
	_ = transferStopCmd.MarkFlagRequired("id")
}

var transferStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the transfer with the given ID.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return printer.Transfer().Stop(id)
	},
}
