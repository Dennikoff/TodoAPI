package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (r *responseWriter) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}
