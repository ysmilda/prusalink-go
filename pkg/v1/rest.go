package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ysmilda/prusalink-go/pkg/rest"
)

// Parses a given byte slice as JSON into type [T] and returns the result.
func parseAsJSON[T any](body []byte, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil //nolint: nilnil
	}
	v := new(T)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// get sends a GET request to the given path and returns the response body.
func (p Printer) get(path string) ([]byte, error) {
	return p.request(http.MethodGet, path, nil, nil)
}

// post sends a POST request to the given path and returns the response body.
func (p Printer) post(path string, data []byte) ([]byte, error) {
	return p.request(http.MethodPost, path, data, nil)
}

// put sends a PUT request to the given path with the given data and headers and returns the response body.
func (p Printer) put(path string, data []byte, headers map[string]string) ([]byte, error) {
	return p.request(http.MethodPut, path, data, headers)
}

// delete sends a DELETE request to the given path and returns the response body.
func (p Printer) delete(path string, headers map[string]string) ([]byte, error) {
	return p.request(http.MethodDelete, path, nil, headers)
}

// patch sends a PATCH request to the given path with the given data and headers and returns the response body.
func (p Printer) patch(path string, data []byte) ([]byte, error) {
	return p.request(http.MethodPatch, path, data, nil)
}

// request sends a request with the given method, path, data, and headers and returns the response body.
func (p Printer) request(method string, path string, data []byte, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = p.headers
	} else {
		for k, v := range p.headers {
			headers[k] = v
		}
	}
	response, err := rest.Request(p.host+path, method, data, headers)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
