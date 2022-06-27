package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang-http-server/app/modules"
)

func Routes(){
	router := mux.NewRouter()
	router.HandleFunc("/", modules.ReturnString)
	router.HandleFunc("/myArray", modules.ReturnArray)
	router.HandleFunc("/getDataById/{id}", modules.ReturnDatabyID)
	log.Fatal(http.ListenAndServe(":8089", router))
}
