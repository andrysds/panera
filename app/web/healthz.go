package web

import (
	"fmt"
	"net/http"
)

func (w *Web) Healthz(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "HEALTHZ OK")
}
