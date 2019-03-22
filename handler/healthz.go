package handler

import (
	"io"
	"net/http"
)

// Healthz is handler function for /healthz
func Healthz(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "OK!\n")
}
