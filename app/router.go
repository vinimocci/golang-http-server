package app

import (
	"net/http"
	"log"
	"fmt"
)

func Routes(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8089", nil))

}

func homePage(w http.ResponseWriter, r *http.Request){	
	fmt.Fprintf(w, "teste")	
 }