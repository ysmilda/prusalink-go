package v1

import (
	"fmt"
)

var (
	ErrNonGcodeFile = fmt.Errorf("file must have .gcode extension")
	ErrEmptyFile    = fmt.Errorf("empty files are not supported")

	ErrStorageNotFound = fmt.Errorf("storage not found")
	ErrAlreadyExists   = fmt.Errorf("already exists")
)

type Printer struct {
	host    string
	key     string
	headers map[string]string
}

func NewPrinter(host string, key string) *Printer {
	return &Printer{
		host: host,
		key:  key,
		headers: map[string]string{
			"X-Api-Key": key,
		},
	}
}

func (p Printer) Version() (*Version, error) {
	return parseAsJSON[Version](p.get("/api/version"))
}

func (p Printer) Info() (*Info, error) {
	return parseAsJSON[Info](p.get("/api/v1/info"))
}

func (p Printer) Status() (*Status, error) {
	return parseAsJSON[Status](p.get("/api/v1/status"))
}

func (p Printer) Storage() (*StorageInfo, error) {
	return parseAsJSON[StorageInfo](p.get("/api/v1/storage"))
}

func (p Printer) Job() *jobHandler {
	return &jobHandler{&p}
}

func (p Printer) Transfer() *transferHandler {
	return &transferHandler{&p}
}

func (p Printer) Files() *filesHandler {
	return &filesHandler{&p}
}

// Warning: This function has not been tested against a printer due to lack of hardware.
func (p Printer) Camera() *cameraHandler {
	return &cameraHandler{&p}
}
