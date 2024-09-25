package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
	v1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

var conn *v1.V1

var v1Cmd = &cobra.Command{
	Use:   "v1",
	Short: "v1 API commands",
	PreRun: func(_ *cobra.Command, _ []string) {
		conn = cli.Printer.V1()
	},
}

func init() {
	cli.RootCmd.AddCommand(v1Cmd)
}
