package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type createPoolResponse struct{
    Uuid string `json:"uuid"`
}

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
    Uuid := uuid.New().String()
    responseObj := &createPoolResponse{
        Uuid: Uuid,
    }
    res, _ := json.Marshal(responseObj)
    CreatePool(Uuid)
    w.WriteHeader(200)
    w.Write(res)
}


type getPoolHandler struct {}
func (h * getPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    
    w.WriteHeader(200)
    w.Write(TestJson())
}
