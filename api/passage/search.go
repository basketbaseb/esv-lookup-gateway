package handler

import (
    "fmt"
    "log"
    "net/http"
    "github.com/basketbaseb/esv-lookup-gateway/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    esvClient := utils.NewClient()

    resp, err := esvClient.GetSearch(r)
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
