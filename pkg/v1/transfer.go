package v1

import (
	"errors"
	"fmt"

	"github.com/ysmilda/prusalink-go/pkg/rest"
)

type transferHandler struct {
	printer *Printer
}

// Returns info about current transfer.
// If there is no transfer, returns nil.
func (t transferHandler) GetInfo() (*Transfer, error) {
	return parseAsJSON[Transfer](t.printer.get("/api/v1/transfer"))
}

// Stops the transfer with the given ID.
func (t transferHandler) Stop(id int) error {
	_, err := t.printer.delete(fmt.Sprintf("/api/v1/transfer/%d", id), nil)
	if err != nil {
		return fmt.Errorf("could not stop transfer: %w", t.parseError(err, id))
	}
	return nil
}

func (t transferHandler) parseError(err error, id int) error {
	if errors.Is(err, rest.ErrNotFound) {
		return fmt.Errorf("transfer with ID %d not found", id)
	}
	return err
}
