package v0

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieves version information about the printer.",
	RunE: func(_ *cobra.Command, _ []string) error {
		version, err := printer.Version()
		if err != nil {
			return err
		}
		return cli.Print(version)
	},
}

func init() {
	v0Cmd.AddCommand(versionCmd)
}
