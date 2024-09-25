package v1

import (
	"github.com/ysmilda/prusalink-go/pkg/printer"
	"github.com/ysmilda/prusalink-go/pkg/utils"
)

type V1 struct {
	conn *printer.Conn
}

func New(conn *printer.Conn) *V1 {
	return &V1{conn}
}

func (v V1) Info() (*Info, error) {
	return utils.ParseAsJSON[Info](v.conn.Get("/api/v1/info"))
}

func (v V1) Status() (*Status, error) {
	return utils.ParseAsJSON[Status](v.conn.Get("/api/v1/status"))
}

func (v V1) Storage() (*Storage, error) {
	return utils.ParseAsJSON[Storage](v.conn.Get("/api/v1/storage"))
}

func (v V1) Transfer() (*Transfer, error) {
	return utils.ParseAsJSON[Transfer](v.conn.Get("/api/v1/transfer"))
}

func (v V1) Job() *jobHandler {
	return &jobHandler{v.conn}
}

func (v V1) Files() *filesHandler {
	return &filesHandler{v.conn}
}
