package main

import (
    "fmt"
    "bytes"
    "encoding/json"
    "log"
    "net/http"
)

type TransformedRequest struct {
    Event          string            
    EventType      string            
    AppID          string            
    UserID         string            
    MessageID      string            
    PageTitle      string            
    PageURL        string            
    BrowserLanguage string           
    ScreenSize     string            
    Attributes     map[string]Field  
    Traits         map[string]Field  
}

type Field struct {
    Value string 
    Type  string 
}

func RequestConverted(req jsonRequest) TransformedRequest {
    attributes := map[string]Field{
        req.Atrk1: {Value: req.Atrv1, Type: req.Atrt1},
        req.Atrk2: {Value: req.Atrv2, Type: req.Atrt2},
    }

    traits := map[string]Field{
        req.Uatrk1: {Value: req.Uatrv1, Type: req.Uatrt1},
        req.Uatrk2: {Value: req.Uatrv2, Type: req.Uatrt2},
        req.Uatrk3: {Value: req.Uatrv3, Type: req.Uatrt3},
    }

    return TransformedRequest{
        Event:          req.Ev,
        EventType:      req.Et,
        AppID:          req.Id,
        UserID:         req.Uid,
        MessageID:      req.Mid,
        PageTitle:      req.T,
        PageURL:        req.P,
        BrowserLanguage: req.L,
        ScreenSize:     req.Sc,
        Attributes:     attributes,
        Traits:         traits,
    }
}

func worker() {
    for req := range requestChannel {
        transformedReq := RequestConverted(req)
        jsonData, err := json.Marshal(transformedReq)
        fmt.Println(bytes.NewBuffer(jsonData))
        if err != nil {
            log.Println("Error marshaling JSON:", err)
            continue
        }
        _, err = http.Post("https://webhook.site/3f3cdefc-fe7c-4180-8a36-6eaa8c7d6aff", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            log.Println("Error sending request to webhook:", err)
        }
    }
}
