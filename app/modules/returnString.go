package modules

import (
	"fmt"
	"net/http"
)

func ReturnString(w http.ResponseWriter, r *http.Request){	
	fmt.Fprintf(w, "testing some data as response =D")	
 }