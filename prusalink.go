package prusalink

import (
	"github.com/ysmilda/prusalink-go/pkg/printer"
	"github.com/ysmilda/prusalink-go/pkg/utils"
	v1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

type Printer struct {
	*printer.Conn
}

func NewPrinter(host string, key string) *Printer {
	return &Printer{printer.NewConn(host, key)}
}

func (p Printer) Version() (*Version, error) {
	return utils.ParseAsJSON[Version](p.Get("/api/version"))
}

func (p *Printer) V1() *v1.V1 {
	return v1.New(p.Conn)
}

type Version struct {
	API            string       `json:"api"`
	Server         string       `json:"server"`
	NozzleDiameter float64      `json:"nozzle_diameter"`
	Text           string       `json:"text"`
	Hostname       string       `json:"hostname"`
	Capabilities   Capabilities `json:"capabilities,omitempty"` // Additional capabilities the printer has.
}

type Capabilities struct {
	UploadByPUT bool `json:"upload-by-put"` // The printer supports uploading GCodes by the PUT method.
}
