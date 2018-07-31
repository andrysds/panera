package main

import (
	"net/http"
	"os"

	"github.com/andrysds/panera/app"
)

func main() {
	port := os.Getenv("PORT")
	go http.ListenAndServe(":"+port, nil)

	panera := app.NewPanera()
	panera.Run()
}
