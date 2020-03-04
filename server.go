package main

import (
    "fmt"
    "log"
    "net/http"
    "./utils"
)

func GetPassageHtml(w http.ResponseWriter, r *http.Request) {
    esvClient := utils.NewClient()

    resp, err := esvClient.GetHTML(r)
    if err != nil {
        log.Fatalln(err)
        return
    }

    defer resp.Body.Close()

    body, err := utils.ParseResponse(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // Send response to client
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(body))
}

func GetAudio(w http.ResponseWriter, r *http.Request) {
    esvClient := utils.NewClient()

    resp, err := esvClient.GetAudio(r)
    if err != nil {
        log.Fatalln(err)
        return
    }

    defer resp.Body.Close()

    body, err := utils.ParseResponse(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // Send response to client
    w.Header().Set("Content-Type", "audio/mpeg")
    fmt.Fprint(w, string(body))
}

func main() {
    http.HandleFunc("/api/passage/html", GetPassageHtml)
    http.HandleFunc("/api/passage/audio", GetAudio)
    http.ListenAndServe(":8090", nil)
}
