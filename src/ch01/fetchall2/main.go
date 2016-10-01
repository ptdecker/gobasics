//Fethall fetches URLs in parallel and reports their times and sizes
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    fetchall()
    fetchall()
}

func fetchall() {

    start := time.Now()
    ch    := make(chan string)

    // Fetch URLs in parallel

    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }

    // Receive results from each channel

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

    start := time.Now()
    
    // Attempt to open a connection to a URL and write out any errors
    // to the channel

    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    // Read the body of the page, but discard the contents keeping only
    // the size of the body and any errors. Write errors out to the 
    // channel

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    // Write out time, size, and the URL read to the channel

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

