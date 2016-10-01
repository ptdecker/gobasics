//Server2 is a minimal "echo" and counter server
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu    sync.Mutex
var count int

func main() {

    // Setup route handlers

    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)

    // Start server

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echos the path component of the requested URL
// and tracks how often it has been called

func handler(w http.ResponseWriter, r *http.Request) {

    // Increment global call counter

    mu.Lock()
    count++
    mu.Unlock()

    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Handler echos the number of calls received so far.  Calls
// to this handler do not increment call count

func counter(w http.ResponseWriter, r *http.Request) {

    mu.Lock()
    fmt.Fprintf(w, "Count: %d\n", count)
    mu.Unlock()
}
