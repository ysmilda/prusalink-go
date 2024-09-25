package v1_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/prusalink-go/pkg/printer"
	v1 "github.com/ysmilda/prusalink-go/pkg/v1"
)

var p *v1.V1

func init() {
	host := os.Getenv("PRUSALINK_HOST")
	key := os.Getenv("PRUSALINK_KEY")
	if host == "" || key == "" {
		panic("PRUSALINK_HOST and PRUSALINK_KEY must be set")
	}
	p = v1.New(printer.NewConn(host, key))
}

func TestInfo(t *testing.T) {
	_, err := p.Info()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStatus(t *testing.T) {
	status, err := p.Status()
	if err != nil {
		t.Fatal(err)
	}

	// The storage path and read-only status are the only stable values.
	assert.Equal(t, "/usb/", status.Storage.Path)
	assert.Equal(t, false, status.Storage.ReadOnly)
}

func TestGetStorage(t *testing.T) {
	storage, err := p.Storage()
	if err != nil {
		t.Fatal(err)
	}

	// The USB storage is the only stable value.
	assert.Contains(t, storage.StorageList, v1.StorageInfo{
		Name:      "usb",
		Type:      "USB",
		Path:      "/usb/",
		Available: true,
		ReadOnly:  false,
	})
}
