package main

import "net/http"

func handleerr(w http.ResponseWriter,r *http.Request){
	respondError(w,200,"something went wrong")
}