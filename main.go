package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonRequest struct {
    Ev   string 
    Et   string 
    Id   string 
    Uid  string 
    Mid  string 
    T    string 
    P    string 
    L    string 
    Sc   string 
    Atrk1 string
    Atrv1 string
    Atrt1 string
    Atrk2 string
    Atrv2 string
    Atrt2 string
    Uatrk1 string
    Uatrv1 string
    Uatrt1 string
    Uatrk2 string
    Uatrv2 string
    Uatrt2 string
    Uatrk3 string
    Uatrv3 string
    Uatrt3 string
}

var requestChannel = make(chan jsonRequest)

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
        return
    }

    var req jsonRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    requestChannel <- req
    fmt.Fprint(w, "Success")
}

func main() {
    http.HandleFunc("/", handler)
    go worker()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
