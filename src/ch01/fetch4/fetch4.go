// Fetch4 prints the content found at each specified URL using io.Copy to avoid the
// requirement of a buffer. It also adds 'http://' if missing from the passed URL.
// And, prints the status code as well.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {

    for _, url := range os.Args[1:] {

        url = strings.ToLower(url)
        if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
            url = "http://" +  url
        }

        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        written, err := io.Copy(os.Stdout, resp.Body)
        fmt.Println(resp.Status)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: while copyiing from %s to stdout (%u written): %v\n", url, written, err)
            os.Exit(1)
        }

    }
}
