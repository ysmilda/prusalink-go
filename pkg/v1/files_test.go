package v1_test

import (
	"errors"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/prusalink-go/pkg/printer"
)

// TestFiles will test the files API.
//
// After a file upload the printer will show that file on the display. When we than go to delete the file it will
// report an error. This means that after running this test the printer needs to have manual intervention to clear
// the error.
func TestFiles(t *testing.T) {
	content, err := os.ReadFile("testdata/benchy.bgcode")
	if err != nil {
		t.Fatal("failed to read testdata/benchy.bgcode:", err)
	}

	// Create a folder
	err = p.Files().CreateFolder("usb", "/test")
	if err != nil && !errors.Is(err, printer.ErrConflict) {
		t.Fatal("unable to create folder", err)
	}

	// Check the folder
	info, err := p.Files().List("usb", "/test")
	if err != nil {
		t.Fatal("unable to list files:", err)
	}

	assert.Equal(t, "TEST", info.Name)
	assert.Equal(t, true, info.IsDir())
	assert.Equal(t, 0, len(info.Children))

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		// Upload a file
		_ = p.Files().Upload("usb", "/test", "benchy.bgcode", content, true, false)
		wg.Done()
	}()

	time.Sleep(1 * time.Second)

	// Check the transfer status
	transfer, err := p.Transfer()
	if err != nil {
		t.Fatal("unable to get transfer status:", err)
	}

	assert.Equal(t, "FROM_CLIENT", transfer.Type)
	assert.Equal(t, "benchy.bgcode", transfer.DisplayName)
	assert.Equal(t, "/usb/test", transfer.Path)
	assert.Equal(t, len(content), transfer.Size)

	wg.Wait()

	// Check the file
	info, err = p.Files().List("usb", "/test/benchy.bgcode")
	if err != nil {
		t.Fatal("unable to list files:", err)
	}

	assert.Equal(t, "benchy.bgcode", *info.DisplayName)
	assert.Equal(t, len(content), *info.Size)
	assert.Equal(t, "PRINT_FILE", info.Type)
	assert.Equal(t, false, info.IsDir())

	// Delete the file
	err = p.Files().Delete("usb", "/test/benchy.bgcode", false)
	if err != nil {
		t.Fatal("unable to delete file:", err)
	}

	// Delete the folder
	err = p.Files().Delete("usb", "/test", false)
	if err != nil {
		t.Fatal("unable to delete folder:", err)
	}
}
