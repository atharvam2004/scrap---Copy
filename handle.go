package main

import "net/http"

func handle(w http.ResponseWriter,r *http.Request){
	respondWithJson(w,200,struct{}{})
}