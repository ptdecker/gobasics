//Norespserver is a minimal server that does not respond to requests.  Use it as
//a testing tool for Exercise 1.11
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echos the path component of the requested URL
// and then hangs in an infinte loop

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    fmt.Println("Going off the deep end . . .\n")
    for {
        // Do nothing forever
    }
}

