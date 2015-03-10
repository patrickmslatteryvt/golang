package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

const (
  port = ":8080"
)

var calls = 0

type Response map[string]interface{}

func (r Response) String() (s string) {
    b, err := json.Marshal(r)
    if err != nil {
        s = ""
        return
    }
    s = string(b)
    return
}

func Sorry(w http.ResponseWriter, r *http.Request) {
  calls++
  w.Header().Set("Content-Type", "application/json")
  // http://blog.golang.org/json-and-go
  // sorry := {"message": "Sorry, I do not recognize that authorization header token.", "error": true, "calls": "This API has been called times."}
  // json.NewEncoder(w).Encode(sorry)
  json.NewEncoder(w).Encode( Response{"error": true, "message": "Sorry, I do not recognize that authorization header token.", "calls": "This API has been called %d times."})
}

func init() {
    fmt.Printf("Started Sorry server at http://localhost%v.\n", port)
    http.HandleFunc("/", Sorry)
    http.ListenAndServe(port, nil)
}

func main() {}

