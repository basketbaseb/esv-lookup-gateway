package passagehtml

import (
    "fmt"
    "log"
    "os"
    "io/ioutil"
    "net/http"
    "time"
)

func GetPassageHtml(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
  
    // Configure http client with 5 second request timeout
    timeout := time.Duration(5 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }
    
    // Parse query parameters from request
    query, ok := r.URL.Query()["q"]
    if !ok || len(query[0]) < 1 {
        http.Error(w, "/?q= is required.", http.StatusBadRequest)
        return
    }
    
    // Create new request to ESV API and forward query parameters
    request, err := http.NewRequest("GET", "https://api.esv.org/v3/passage/html?q=" + query[0], nil)
    // Add ESV API token to request headers
    apiKey := os.Getenv("ESV_API_TOKEN")
    request.Header.Set("Authorization", "Token " + apiKey)
    
    if err != nil {
        log.Fatalln(err)
    }
    
    // Send the request
    resp, err := client.Do(request)
    if err != nil {
        log.Fatalln(err)
    }
    
    defer resp.Body.Close()
    
    // Parse the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    
    // Send response to client
    fmt.Fprintf(w, string(body))
}