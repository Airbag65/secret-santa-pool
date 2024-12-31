package main

import (
	"net/http"
)

func main(){
    mux := http.NewServeMux()
    
    mux.Handle("/", &homeHandler{})
    mux.Handle("/createpool", &createPoolHandler{})
    mux.Handle("/get", &getPoolHandler{})

    http.ListenAndServe(":8080", mux)
}
