package cli

import "github.com/spf13/cobra"

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

func init() {
	RootCmd.AddCommand(versionCmd)
}
