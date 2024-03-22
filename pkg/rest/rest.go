package rest

import (
	"bytes"
	"fmt"
	"net/http"
)

func Request(url string, method string, body []byte, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
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

	return response, nil
}

var (
	ErrNotModified          = fmt.Errorf("not modified")
	ErrBadRequest           = fmt.Errorf("bad request")
	ErrUnauthorized         = fmt.Errorf("unauthorized")
	ErrForbidden            = fmt.Errorf("forbidden")
	ErrNotFound             = fmt.Errorf("not found")
	ErrRequestTimeout       = fmt.Errorf("request timeout")
	ErrConflict             = fmt.Errorf("conflict")
	ErrUnsupportedMediaType = fmt.Errorf("unsupported media type")
	ErrInternalServerError  = fmt.Errorf("internal server error")
	ErrNotImplemented       = fmt.Errorf("not implemented")
	ErrServiceUnavailable   = fmt.Errorf("service unavailable")
)

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
