package utils

import (
    "os"
    "io"
    "io/ioutil"
    "net/url"
    "net/http"
)

const (
    defaultBaseUrl = "https://api.esv.org"
)

type Client struct {
    BaseURL    *url.URL
    client *http.Client
}

func NewClient() *Client {
    httpClient := http.DefaultClient
    baseURL, _ := url.Parse(defaultBaseUrl)
    c := &Client{
        client: httpClient,
        BaseURL: baseURL,
    }

    return c;
}

// Parse query parameters from request
func ParseQuery(r *http.Request) (string){
    query := r.URL.RawQuery;
    return query
}

// Parse the response
func ParseResponse(r io.Reader) ([]byte, error) {
    body, err := ioutil.ReadAll(r)
    if err != nil {
        return nil, err
    }

    return body, nil
}

// Create new request
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
    u, err := c.BaseURL.Parse(urlStr)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(method, u.String(), nil)
    if err != nil {
        return nil, err
    }

    // Add ESV API token to request headers
    apiKey := os.Getenv("ESV_API_TOKEN")
    req.Header.Set("Authorization", "Token " + apiKey)
    return req, nil
}

// Send the request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

func (c *Client) GetHTML(r *http.Request) (*http.Response, error) {
    // Forward any query parameters
    query := ParseQuery(r)

    req, err := c.NewRequest("GET", "v3/passage/html?" + query)
    if err != nil {
        return nil, err
    }

    resp, err := c.Do(req)
    if err != nil {
        return nil, err
    }

    return resp, nil
}
