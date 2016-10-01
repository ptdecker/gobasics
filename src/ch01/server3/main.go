//Server3 is a minimal "echo" and counter server that returns
//the headers and form data it receives
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

    // Print the call header, host, and remote address

    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

    // Parse and print out any form fields

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }

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
