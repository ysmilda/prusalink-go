package v0

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
	apiv0 "github.com/ysmilda/prusalink-go/pkg/v0"
)

var (
	host    string
	key     string
	printer *apiv0.Printer
)

var v0Cmd = &cobra.Command{
	Use:   "v0",
	Short: "v0 API commands",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		printer = apiv0.NewPrinter(host, key)
	},
}

func init() {
	cli.RootCmd.AddCommand(v0Cmd)

	v0Cmd.PersistentFlags().StringVar(&host, "host", "http://localhost:80", "The host of the PrusaLink server.")
	v0Cmd.PersistentFlags().StringVar(&key, "key", "", "The API key to authenticate with the PrusaLink server.")
	v0Cmd.MarkPersistentFlagRequired("host")
	v0Cmd.MarkPersistentFlagRequired("key")
}
