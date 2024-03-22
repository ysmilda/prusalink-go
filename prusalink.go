package prusalink

import (
	v0 "github.com/ysmilda/prusalink-go/pkg/v0"
	v1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

func NewPrinterV0(host string, key string) *v0.Printer {
	return v0.NewPrinter(host, key)
}

func NewPrinterV1(host string, key string) *v1.Printer {
	return v1.NewPrinter(host, key)
}
