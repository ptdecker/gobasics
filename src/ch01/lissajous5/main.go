// Lissajous5 generates GIF animations of random Lissajous figures with
// interesting rainbow colors. And, includes a built-in web server.  Also
// supports a 'cycles' parameter passed in the URL to set the number of
// cycles:  For example, "http://localhost:8000/?cycles=20"
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "os"
    "strconv"
    "time"
)

var palette = []color.Color{
    color.Black,
    color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
    color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // orange
    color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // yellow
    color.RGBA{0x00, 0x80, 0x00, 0xFF}, // green
    color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
    color.RGBA{0x8A, 0x2B, 0xE2, 0xFF}, // blue-violet
    color.RGBA{0xEE, 0x82, 0xEE, 0xFF}, // violet
    color.RGBA{0xD0, 0x20, 0x90, 0xFF}, // red-violet
}

const (
    blackIndex      = 0 // first color in palette
    redIndex        = 1 // next color in palette
    orangeIndex     = 2 // etc....
    yellowIndex     = 3
    greenIndex      = 4
    blueIndex       = 5
    blueVioletIndex = 6
    violetIndex     = 7
    redVioletIndex  = 8
)

const defaultCycles = 5

func main() {

    // Seed random number generator

    rand.Seed(time.Now().UTC().UnixNano()) // see errata 

    // If called with "web" as a parameter, launch web server
    // otherwise output a GIF

    if len(os.Args) > 1 && os.Args[1] == "web" {
        webServer()
    } else {
        lissajous(os.Stdout, defaultCycles)
    }
}

func webServer() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        if err := r.ParseForm(); err != nil {
            log.Print(err)
        }
        if i, err := strconv.Atoi(r.FormValue("cycles")); i == 0 || err != nil {
            lissajous(w, defaultCycles)
        } else {
            lissajous(w, i)
        }
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {

    const (
        res     = 0.001  // angular resolution
        size    = 100    // image canvas covers [-size..+size]
        nframes = 128    // number of animation frames
        delay   = 8      // delay between frames in 10ms units
    )

    freq  := rand.Float64() * 3.0   // relative frequency of y ocillator
    anim  := gif.GIF{LoopCount: nframes}
    phase := 0.0  // phase difference

    for i := 0; i < nframes; i++ {

        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img  := image.NewPaletted(rect, palette)

        for t := 0.0; t < float64(cycles) * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), uint8(i % redVioletIndex) + 1)
        }

        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }

    gif.EncodeAll(out, &anim) // Encoding errors are ignored
}
