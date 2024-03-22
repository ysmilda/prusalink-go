package v0

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ysmilda/prusalink-go/pkg/rest"
)

func getAndParseAsJSON[T any](p Printer, path string, v *T) (*T, error) {
	body, err := p.get(path)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (p Printer) get(path string) ([]byte, error) {
	response, err := rest.Request(p.host+path, http.MethodGet, nil, p.headers)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Print the response body, just for testing
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	println(string(body))
	return body, nil
}

func (p Printer) put(path string) error {
	response, err := rest.Request(p.host+path, http.MethodPut, nil, p.headers)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Print the response body, just for testing
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	println(string(body))
	return nil
}

func (p Printer) delete(path string) error {
	response, err := rest.Request(p.host+path, http.MethodDelete, nil, p.headers)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Print the response body, just for testing
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	println(string(body))
	return nil
}
