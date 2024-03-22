package v1

import (
	"errors"
	"fmt"

	"github.com/ysmilda/prusalink-go/pkg/rest"
)

type jobHandler struct {
	printer *Printer
}

// Returns info about the current job.
func (j jobHandler) GetCurrent() (*Job, error) {
	return parseAsJSON[Job](j.printer.get("/api/v1/job"))
}

// Stops the job with the given ID.
func (j jobHandler) Stop(id int) error {
	_, err := j.printer.delete(fmt.Sprintf("/api/v1/job/%d", id), nil)
	if err != nil {
		return fmt.Errorf("could not stop job: %w", j.parseError(err, id))
	}
	return nil
}

// Pauses the job with the given ID.
func (j jobHandler) Pause(id int) error {
	_, err := j.printer.put(fmt.Sprintf("/api/v1/job/%d/pause", id), nil, nil)
	if err != nil {
		return fmt.Errorf("could not pause job: %w", j.parseError(err, id))
	}
	return nil
}

// Resumes the job with the given ID.
func (j jobHandler) Resume(id int) error {
	_, err := j.printer.put(fmt.Sprintf("/api/v1/job/%d/resume", id), nil, nil)
	if err != nil {
		return fmt.Errorf("could not resume job: %w", j.parseError(err, id))
	}
	return nil
}

// Continues the job with the given ID.
func (j jobHandler) Continue(id int) error {
	_, err := j.printer.put(fmt.Sprintf("/api/v1/job/%d/continue", id), nil, nil)
	if err != nil {
		return fmt.Errorf("could not continue job: %w", j.parseError(err, id))
	}
	return nil
}

func (j jobHandler) parseError(err error, id int) error {
	if errors.Is(err, rest.ErrNotFound) {
		err = fmt.Errorf("job with ID %d not found", id)
	}
	return err
}
