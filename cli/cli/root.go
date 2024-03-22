package cli

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var PrettyPrint bool

var RootCmd = &cobra.Command{
	Use:   "prusalink",
	Short: "A CLI to interact with a Prusa 3D printer via PrusaLink",
	Long: `A CLI to interact with a Prusa 3D printer via PrusaLink.
	
	PrusaLink is a REST API for Prusa 3D printers. This CLI provides a simple
	interface to interact with the API. It is intended to be used for scripting
	and automation.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&PrettyPrint, "pretty", "p", false, "Pretty print the output.")
}

func Print(in interface{}) error {
	var (
		output []byte
		err    error
	)
	if PrettyPrint {
		output, err = json.MarshalIndent(in, "", "  ")
	} else {
		output, err = json.Marshal(in)
	}
	if err != nil {
		return err
	}
	println(string(output))
	return nil
}
