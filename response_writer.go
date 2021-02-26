package nrfiber

import (
	"net/http"
)

// ResponseWriter imitates http.ResponseWriter
type ResponseWriter struct {
	statusCode int
	header     http.Header
	body       []byte
}

// Header implementation
func (rw *ResponseWriter) Header() http.Header {
	if rw.header == nil {
		rw.header = make(http.Header)
	}
	return rw.header
}

// WriteHeader implementation
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

// Write implementation
func (rw *ResponseWriter) Write(p []byte) (int, error) {
	rw.body = append(rw.body, p...)
	return len(p), nil
}
