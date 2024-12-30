package main

import (
	"net/http"
)

func main(){
    mux := http.NewServeMux()
    
    mux.Handle("/", &homeHandler{})
    mux.Handle("/create", &createPoolHandler{})
    mux.Handle("/get", &getPoolHandler{})

    http.ListenAndServe(":8080", mux)
}


