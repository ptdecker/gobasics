// Mandelbrot7 provides a web service that generates colorized PNG image of the Mandelbrot fractal
// and supports x, y, and zoom values as parameters passed in the HTTP request

// c.f. http://www.fractalforums.com/programming/newbie-how-to-map-colors-in-the-mandelbrot-set/
// c.f. https://en.wikipedia.org/wiki/Mandelbrot_set
// c.f. "The Science of Fractal Images", Springer-Verlang, ISBN 0-387-96608-0

package main

import (
    "image"
    "image/color"
    "image/png"
    "io"
    "log"
    "math"
    "math/cmplx"
    "net/http"
    "strconv"
)

func main() {

    handler := func(w http.ResponseWriter, r *http.Request) {
        var x, y, z float64
        var err error
        if err := r.ParseForm(); err != nil {
            log.Print(err)
        }
        if z, err = strconv.ParseFloat(r.FormValue("zoom"), 64); z == 0 || err != nil {
            z = 1
        }
        if x, err = strconv.ParseFloat(r.FormValue("x"), 64); err != nil {
            x = 0
        }
        if y, err = strconv.ParseFloat(r.FormValue("y"), 64); err != nil {
            y = 0
        }
        genPNG(w, -2 / z + x, -2 / z + y, 2 / z + x, 2 / z + y)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// Generate Mandelbrot PNG
func genPNG(out io.Writer, xmin float64, ymin float64, xmax float64, ymax float64) {

    const (
        width = 1024
        height = 1024
    )

    // Iterate through all pixels in the image mapping each pixel into Mandelbrot space
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(out, img) // NOTE: ignores errors
}

// Calculate number of iterations required to escape a circle radius of '2' 
func mandelbrot(z complex128) color.Color {

    const iterations = 200

    var v complex128

    for n := uint8(0); n < iterations; n++ {
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            return map2RGB(255 - 16 * n)
        }
    }

    return color.Black
}

// Map the number of iterations into the color pallet
//
// c.f. http://www.rapidtables.com/convert/color/hsv-to-rgb.htm

func map2RGB(n uint8) (color.Color) {

    h := 360 * float64(n) / 256
    x := uint8(255 * (1 - math.Abs(math.Mod(h / 60, 2) - 1)))

    switch {
        case h <  60:
            return color.RGBA{ 255, x,  0, 255 }
        case h < 120:
            return color.RGBA{ x, 255,  0, 255 }
        case h < 180:
            return color.RGBA{  0, 255, x, 255 }
        case h < 240:
            return color.RGBA{  0, x, 255, 255 }
        case h < 300:
            return color.RGBA{ x,  0, 255, 255 }
        default:
            return color.RGBA{ 255,  0, x, 255 }
    }

    return color.Black
}
