// Mandelbrot5 re-implements Mandelbrot1 using big.Float to emit colorized PNG images of the Mandelbrot fractal
// and includes benchmarking

// c.f. http://www.fractalforums.com/programming/newbie-how-to-map-colors-in-the-mandelbrot-set/
// c.f. https://en.wikipedia.org/wiki/Mandelbrot_set
// c.f. "The Science of Fractal Images", Springer-Verlang, ISBN 0-387-96608-0

package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/big"
    "os"
)

func main() {

    const (
        xmin, ymin, xmax, ymax  = -2, -2, 2, 2  // Mandelbrot space coordinates
        width, height           = 1024, 1024    // Rendered image size in pixels
    )

    // Iterate through all pixels in the image mapping each pixel into Mandelbrot space

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := int64(0); py < height; py++ {
        y := new(big.Float).SetInt64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := new(big.Float).SetInt(px) / width * (xmax - xmin) + xmin
            img.Set(px, py, mandelbrot(x, y)) // where a + xi represents the imaginary number
        }
    }
    png.Encode(os.Stdout, img) // NOTE: ignores errors

}

// Calculate number of iterations required to escape a circle radius of '2' 
// r = real part, i = imaginary part
func mandelbrot(r Float, i Float) color.Color {

    const iterations = 200

    vr := New(big.Float) // real part
    vi := New(big.Float) // imaginary part
    t  := New(big.Float)
    x  := New(big.Float) // temp variable
    y  := New(big.Float)
    z  := New(big.Float)

    for n := uint8(0); n < iterations; n++ {
        vr.Add(x.Sub(y.Mul(vr, vr), z.Mul(vi, vi)), r) 
        vi.Add(x.Sub(y.Mul(vi, vr), z.Mul(vr, vi)), i)
        if t.Add(x.Add(y.Mul(vr, vr), z.Mul(vi, vi))) > 2 {
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
