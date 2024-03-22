package main

import (
	"github.com/ysmilda/prusalink-go/cli/cli"
	_ "github.com/ysmilda/prusalink-go/cli/cli/v0"
	_ "github.com/ysmilda/prusalink-go/cli/cli/v1"
)

func main() {
	cli.Execute()
}
