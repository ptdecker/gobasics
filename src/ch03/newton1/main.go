// Mandelbrot emits colorized PNG image of the Mandelbrot fractal

// c.f. http://www.fractalforums.com/programming/newbie-how-to-map-colors-in-the-mandelbrot-set/
// c.f. https://en.wikipedia.org/wiki/Mandelbrot_set
// c.f. "The Science of Fractal Images", Springer-Verlang, ISBN 0-387-96608-0

package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "os"
)

func main() {

    const (
        xmin, ymin, xmax, ymax  = -2, -2, 2, 2  // Mandelbrot space coordinates
        width, height           = 1024, 1024    // Rendered image size in pixels
    )

    // Iterate through all pixels in the image mapping each pixel into Mandelbrot space

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, newton(z))
        }
    }
    png.Encode(os.Stdout, img) // NOTE: ignores errors

}

// Calculate z^4 - 1 = 0 using Newtons method
//
// c.f. https://github.com/adonovan/gopl.io/blob/master/ch3/mandelbrot/main.go

func newton(z complex128) color.Color {

    const (
        iterations = 37
        contrast = 16
        threshold = 0.000001
    )

    for n := uint8(0); n < iterations; n++ {
        z = z - (z - 1 / (z * z * z)) / 4
        if cmplx.Abs(z * z * z * z - 1) < threshold {
            return map2RGB(255 - contrast * n)
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
