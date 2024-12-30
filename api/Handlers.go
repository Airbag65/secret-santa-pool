package main

import "net/http"


type homeHandler struct{}

func (h *homeHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}


type createPoolHandler struct {}

func (h * createPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}


type getPoolHandler struct {}
func (h * getPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}
