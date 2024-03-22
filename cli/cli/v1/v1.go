package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
	apiv1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

var (
	host    string
	key     string
	printer *apiv1.Printer
)

var v1Cmd = &cobra.Command{
	Use:   "v1",
	Short: "v1 API commands",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		printer = apiv1.NewPrinter(host, key)
	},
}

func init() {
	cli.RootCmd.AddCommand(v1Cmd)

	v1Cmd.PersistentFlags().StringVar(&host, "host", "http://localhost:80", "The host of the PrusaLink server.")
	v1Cmd.PersistentFlags().StringVar(&key, "key", "", "The API key to authenticate with the PrusaLink server.")
	_ = v1Cmd.MarkPersistentFlagRequired("host")
	_ = v1Cmd.MarkPersistentFlagRequired("key")
}
