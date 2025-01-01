package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type createPoolResponse struct{
    Uuid string `json:"uuid"`
}

type addMemberRequest struct{
    Uid string `json:"uid"`
    Person Person `json:"person"`
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

    res, err := json.Marshal(responseObj)
    if err != nil {
        w.WriteHeader(500)
        w.Write([]byte("Something went wrong"))
        return
    }

    if err := CreatePool(Uuid); err != nil {
        w.WriteHeader(500)
        w.Write([]byte("Something went wrong"))
        return
    }
    w.WriteHeader(200)
    w.Write(res)
}

func GetParams(r *http.Request) map[string]string {
    params := strings.Split(r.URL.String(), "?")[1]
    var paramList []string
    vars := map[string]string{}
    if strings.Contains(params, "&"){
        paramList = strings.Split(params, "&")
    } else {
        paramList = []string{params}
    }
    for _, item := range paramList {
        param := strings.Split(item, "=")
        vars[param[0]] = param[1]
    }
    return vars
}

type getPoolHandler struct {}
func (h * getPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }

    // "eb2f08a5-12de-4253-9e14-cad1bf8f762e"
    if len(strings.Split(r.URL.String(), "?")) == 1 {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    vars := GetParams(r)
    pool, err := GetPool(vars["uid"])
    if err != nil{
        if err.Error() == "No such pool" {
            w.WriteHeader(404)
            w.Write([]byte("Pool does not exist"))
        } else if err.Error() == "Error" {
            w.WriteHeader(500)
            w.Write([]byte("Something went wrong"))
        }
        return 
    }
    
    w.WriteHeader(200)
    w.Write(pool)
}

type addPersonHandler struct {}
func (h *addPersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    var req addMemberRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    if r.Header.Get("Content-Type") != "application/json" {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    err = InsertMember(&req)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte("Pool does not exist"))
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}
