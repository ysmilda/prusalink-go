package v1

import (
	"github.com/ysmilda/prusalink-go/pkg/printer"
)

type V1 struct {
	conn *printer.Conn
}

func New(conn *printer.Conn) *V1 {
	return &V1{conn}
}

// Info returns general information about the printer.
func (v V1) Info() (*Info, error) {
	return printer.ParseAsJSON[Info](v.conn.Get("/api/v1/info"))
}

// Status returns condensed information about various printer states.
func (v V1) Status() (*Status, error) {
	return printer.ParseAsJSON[Status](v.conn.Get("/api/v1/status"))
}

// Storage returns information about the available storage.
func (v V1) Storage() (*Storage, error) {
	return printer.ParseAsJSON[Storage](v.conn.Get("/api/v1/storage"))
}

// Transfer returns the current transfer status.
// However it seems that the transfer information in the Status request is updated more frequently.
func (v V1) Transfer() (*Transfer, error) {
	return printer.ParseAsJSON[Transfer](v.conn.Get("/api/v1/transfer"))
}

func (v V1) Job() *jobHandler {
	return &jobHandler{v.conn}
}

func (v V1) Files() *filesHandler {
	return &filesHandler{v.conn}
}
