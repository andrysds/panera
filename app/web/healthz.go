package web

import (
	"fmt"
	"net/http"
)

func (w *Web) HandleHealthz(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "HEALTHZ OK")
}
