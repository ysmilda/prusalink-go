package prusalink_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/prusalink-go"
)

var p *prusalink.Printer

func init() {
	host := os.Getenv("PRUSALINK_HOST")
	key := os.Getenv("PRUSALINK_KEY")
	if host == "" || key == "" {
		panic("PRUSALINK_HOST and PRUSALINK_KEY must be set")
	}
	p = prusalink.NewPrinter(host, key)
}

func TestVersion(t *testing.T) {
	version, err := p.Version()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "2.0.0", version.API)
	assert.Equal(t, "PrusaLink", version.Text)
}
