// Mandelbrot2 emits colorized PNG image of the Mandelbrot fractal using super-sampling to reduce aliasing

// c.f. http://www.fractalforums.com/programming/newbie-how-to-map-colors-in-the-mandelbrot-set/
// c.f. https://en.wikipedia.org/wiki/Mandelbrot_set
// c.f. "The Science of Fractal Images", Springer-Verlang, ISBN 0-387-96608-0
// c.f. https://web.cs.wpi.edu/~matt/courses/cs563/talks/antialiasing/methods.html

package main

import (
    "image"
    "image/color"
    "image/png"
    "io"
    "math"
    "math/cmplx"
    "os"
)

const (
    width  = 2048
    height = 2048
)

var virtualImage [width][height]color.Color
var sampledImage [width / 2][height / 2]color.Color

func main() {

    calcVirtualImage(-2, -2, 2, 2)
    sampleImage()
    renderPNGImage(os.Stdout)
}

// Render an image to a PNG
func renderPNGImage(w io.Writer) {
    img := image.NewRGBA(image.Rect(0, 0, width / 2, height / 2))
    for py := 0; py < (height / 2); py++ {
        for px := 0; px < (width / 2); px++ {
            img.Set(px, py, sampledImage[px][py])
        }
    }
    png.Encode(w, img) // NOTE: ignores errors
}

// Create a new image by supersampling another image
func sampleImage() {
    for py := 0; py < (height / 2); py++ {
        for px := 0; px < (width / 2); px++ {

            r1, g1, b1, _ := color.Color.RGBA(virtualImage[2 * px][2 * py])
            r2, g2, b2, _ := color.Color.RGBA(virtualImage[2 * px + 1][2 * py])
            r3, g3, b3, _ := color.Color.RGBA(virtualImage[2 * px][2 * py + 1])
            r4, g4, b4, _ := color.Color.RGBA(virtualImage[2 * px + 1][2 * py + 1])

            r := (r1 + r2 + r3 + r4) / 4
            g := (g1 + g2 + g3 + g4) / 4
            b := (b1 + b2 + b3 + b4) / 4

            sampledImage[px][py] = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
        }
    }
}

// Calulate a Mandelbrot set
func calcVirtualImage(xmin, ymin, xmax, ymax float64) {
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            virtualImage[px][py] = mandelbrot(z)
        }
    }
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
