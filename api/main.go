package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main(){
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.WriteHeader(200)
        w.Write([]byte("OK"))
    })
    mux.Handle("/createpool", &createPoolHandler{})
    mux.Handle("/get", &getPoolHandler{})
    mux.Handle("/addmember", &addPersonHandler{})
    mux.Handle("/lottery", &performLotteryHandler{})

    handler := cors.Default().Handler(mux)

    http.ListenAndServe(":8080", handler)
}
