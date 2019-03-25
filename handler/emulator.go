package handler

import (
	"net/http"

	"github.com/andrysds/panera/template"
)

// Emulator is handler function for GET /emulator
func Emulator(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	data := struct {
		templateData
		Command  string
		Response string
	}{
		templateData: templateData{
			PageTitle:  "Emulator",
			FormAction: "/emulator",
		},
		Command:  command,
		Response: Message(command),
	}
	template.Execute(w, "emulator.html", data)
}
