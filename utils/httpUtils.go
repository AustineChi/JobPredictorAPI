package utils

import (
    "io"
    "net/http"
)

// HttpGet makes an HTTP GET request to the specified URL and returns the response body
func HttpGet(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Using io.ReadAll to read the response body
    return io.ReadAll(resp.Body)
}
