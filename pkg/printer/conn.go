package printer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Conn struct {
	host    string
	headers map[string]string
}

func NewConn(host string, key string) *Conn {
	return &Conn{
		host: host,
		headers: map[string]string{
			"X-Api-Key": key,
		},
	}
}

// get a GET request to the given path and returns the response body.
func (c Conn) Get(path string) ([]byte, error) {
	return c.request(http.MethodGet, path, nil, nil)
}

// post sends a POST request to the given path and returns the response body.
func (c Conn) Post(path string, data []byte) ([]byte, error) {
	return c.request(http.MethodPost, path, data, nil)
}

// put sends a PUT request to the given path with the given data and headers and returns the response body.
func (c Conn) Put(path string, data []byte, headers map[string]string) ([]byte, error) {
	return c.request(http.MethodPut, path, data, headers)
}

// delete sends a DELETE request to the given path and returns the response body.
func (c Conn) Delete(path string, headers map[string]string) ([]byte, error) {
	return c.request(http.MethodDelete, path, nil, headers)
}

// patch sends a PATCH request to the given path with the given data and headers and returns the response body.
func (c Conn) Patch(path string, data []byte) ([]byte, error) {
	return c.request(http.MethodPatch, path, data, nil)
}

// request sends a request with the given method, path, data, and headers and returns the response body.
func (c Conn) request(method string, path string, data []byte, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	for k, v := range c.headers {
		headers[k] = v
	}

	request, err := http.NewRequest(method, c.host+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if err := responseOK(response); err != nil {
		response.Body.Close()
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func responseOK(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}

	switch response.StatusCode {
	case http.StatusNotModified:
		return ErrNotModified
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusRequestTimeout:
		return ErrRequestTimeout
	case http.StatusConflict:
		return ErrConflict
	case http.StatusUnsupportedMediaType:
		return ErrUnsupportedMediaType
	case http.StatusInternalServerError:
		return ErrInternalServerError
	case http.StatusNotImplemented:
		return ErrNotImplemented
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	default:
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}
