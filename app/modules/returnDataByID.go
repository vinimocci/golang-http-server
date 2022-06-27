package modules

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func ReturnDatabyID(w http.ResponseWriter, r *http.Request){	
	vars := mux.Vars(r)
    sentID := vars["id"]

	if sentID == "1" {
		fmt.Fprintf(w, "Correct ID!")
	} else {
		fmt.Fprintf(w, "please verify your sent ID")
	}
 }
