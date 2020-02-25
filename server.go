package main

import (
    "fmt"
    "log"
    "net/http"
    "./utils"
)

func GetPassageHtml(w http.ResponseWriter, r *http.Request) {
    esvClient := rest.NewClient()
      
    resp, err := esvClient.GetHTML(r)
    if err != nil {
        log.Fatalln(err)
        return
    }
    
    defer resp.Body.Close()
    
    body, err := rest.ParseResponse(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    
    // Send response to client
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(body))
}

func main() {
    http.HandleFunc("/passage/html", GetPassageHtml)
    http.ListenAndServe(":8090", nil)
}