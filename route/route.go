package route

import (
	"net/http"

	"github.com/andrysds/panera/handler"
)

// Init initializes http routes
func Init() {
	http.HandleFunc("/healthz", handler.Healthz)
}
