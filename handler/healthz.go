package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andrysds/panera/entity"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := map[string]string{
		"message": entity.OKMessage,
	}
	json.NewEncoder(w).Encode(result)
}
