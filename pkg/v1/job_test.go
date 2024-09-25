package v1_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

func TestJobs(t *testing.T) {
	// Make sure that the printer is idle
	job, err := p.Job().Get()
	if err != nil && !errors.Is(err, v1.ErrNoJob) {
		t.Fatal("unable to get job:", err)
	}

	if job != nil {
		t.Fatal("A job is running, please cancel it before running this test")
	}

	content, err := os.ReadFile("testdata/noop.gcode")
	if err != nil {
		t.Fatal("failed to read testdata/noop.gcode:", err)
	}

	// Upload a file
	err = p.Files().Upload("usb", "/", "noop.gcode", content, true, true)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	job, err = p.Job().Get()
	if err != nil {
		t.Fatal(err)
	}

	if job == nil {
		t.Fatal("no job found")
	}

	assert.Equal(t, "noop.gcode", job.File.DisplayName)
	assert.Equal(t, "PRINTING", job.State)

	err = p.Job().Pause(job.ID)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	job, err = p.Job().Get()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "PAUSED", job.State)

	// TODO: This doesn't work, the printer doesn't resume the job
	err = p.Job().Resume(job.ID)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	job, err = p.Job().Get()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "PRINTING", job.State)

	err = p.Job().Stop(job.ID)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	job, err = p.Job().Get()
	if err != nil && !errors.Is(err, v1.ErrNoJob) {
		t.Fatal(err)
	}

	if job != nil {
		t.Fatal("job is still running")
	}

	err = p.Files().Delete("usb", "noop.gcode", false)
	if err != nil {
		t.Fatal(err)
	}
}
