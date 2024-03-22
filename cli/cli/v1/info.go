package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Retrieves information about the printer.",
	RunE: func(_ *cobra.Command, _ []string) error {
		info, err := printer.Info()
		if err != nil {
			return err
		}
		return cli.Print(info)
	},
}

func init() {
	v1Cmd.AddCommand(infoCmd)
}
