// Fetch2 prints the content found at each specified URL using io.Copy to avoid the
// requirement of a buffer
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {

    for _, url := range os.Args[1:] {

        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        written, err := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: while copyiing from %s to stdout (%u written): %v\n", url, written, err)
            os.Exit(1)
        }

    }
}
