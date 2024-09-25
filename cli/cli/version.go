package cli

import (
	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
	data    string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieves the version of the CLI.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		version := struct {
			Version string `json:"version"`
			Commit  string `json:"commit"`
			Data    string `json:"data"`
		}{
			Version: version,
			Commit:  commit,
			Data:    data,
		}
		return Print(version)
	},
}

var versionPrinterCmd = &cobra.Command{
	Use:   "versionprinter",
	Short: "Retrieves the version information from the printer.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		version, err := Printer.Version()
		if err != nil {
			return err
		}
		return Print(version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(versionPrinterCmd)
}
