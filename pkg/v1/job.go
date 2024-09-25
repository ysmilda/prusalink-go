package v1

import (
	"errors"
	"fmt"

	"github.com/ysmilda/prusalink-go/pkg/printer"
	"github.com/ysmilda/prusalink-go/pkg/utils"
)

type jobHandler struct {
	conn *printer.Conn
}

// Get retrieves information on the current job.
func (j jobHandler) Get() (*Job, error) {
	return utils.ParseAsJSON[Job](j.conn.Get("/api/v1/job"))
}

// Stop stops the job with the given ID.
func (j jobHandler) Stop(id int) error {
	_, err := j.conn.Delete(fmt.Sprintf("/api/v1/job/%d", id), nil)
	if err != nil {
		return fmt.Errorf("could not stop job: %w", j.parseError(err, id))
	}
	return nil
}

// Pause pauses the job with the given ID.
func (j jobHandler) Pause(id int) error {
	_, err := j.conn.Put(fmt.Sprintf("/api/v1/job/%d/pause", id), nil, nil)
	if err != nil {
		return fmt.Errorf("could not pause job: %w", j.parseError(err, id))
	}
	return nil
}

// Resume resumes the job with the given ID.
func (j jobHandler) Resume(id int) error {
	_, err := j.conn.Put(fmt.Sprintf("/api/v1/job/%d/resume", id), nil, nil)
	if err != nil {
		return fmt.Errorf("could not resume job: %w", j.parseError(err, id))
	}
	return nil
}

func (j jobHandler) parseError(err error, id int) error {
	if errors.Is(err, printer.ErrNotFound) {
		err = fmt.Errorf("job with ID %d not found", id)
	}
	return err
}
