package modules

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func ReturnArray(w http.ResponseWriter, r *http.Request){	
	myArray, _ := json.Marshal([]string{"Vini", "John", "Carl"})
	fmt.Fprintf(w, "%s", myArray)	
 }
