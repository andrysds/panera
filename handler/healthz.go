package handler

import (
	"fmt"
	"net/http"

	"github.com/andrysds/panera/entity"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, entity.OKMessage)
}
