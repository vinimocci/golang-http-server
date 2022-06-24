package app

import (
	"log"
	"net/http"
	"golang-http-server/app/modules"
)

func Routes(){
	http.HandleFunc("/", modules.ReturnString)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

/* 
	Test this one below to return json in Fprintf
	https://stackoverflow.com/questions/14567324/how-to-print-out-a-json-in-go 
*/