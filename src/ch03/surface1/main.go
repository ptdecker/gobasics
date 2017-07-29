// Surface1 computes an SVG rendering of a 3-D surface function.  It also includes guard code to
// ignore any polygons whose corner values are infinte or NaN

// c.f. https://developer.mozilla.org/en-US/docs/Web/SVG
// c.f. https://jwatt.org/svg/authoring/

package main

import (
    "fmt"
    "io"
    "math"
    "os"
)

const (
    width   = 600                  // canvas width in pixels
    height  = 320                  // canvas height in pixels
    cells   = 100                  // number of grid cells
    xyrange = 30.0                 // axis ranges (-xyrange..+xyrange)
    xyscale = width / 2 / xyrange  // pixels per x or y unit
    zscale  = height * 0.4         // pizels per z unit
    angle   = math.Pi / 6          // angle of x and y axis (30 deg)
)

var out io.Writer = os.Stdout // modified during testing

var sin30 = math.Sin(angle)
var cos30 = math.Cos(angle)

// Definition of 3-D surface function to render
func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

// Return coordinates of the corner of a cell on a 2-D canvas
func corner(i, j int) (float64, float64, bool) {
    
    // Find point (x, y) at corner of cell (i,j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute the surface height z using the surface function
    z := f(x, y)

    // Project (x, y, z) isometrically onto the 2-D SVG canvas (sx, sy)
    sx := width  / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return sx, sy, math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)
}

// Generate the SVG
func genSVG() {

    var ax, ay, bx, by, cx, cy, dx, dy float64
    var err bool

    // Open SVG tag
    fmt.Fprintf(out, "<svg " +
        "version='1.1' " +
        "baseProfile='full' " +
        "xmlns='http://www.w3.org/2000/svg' " +
        "xmlns:xlink='http://www.w3.org/1999/xlink' " +
        "xmlns:ev='http://www.w3.org/2001/xml-events' " +
        "stroke='grey' " +
        "stroke-width='0.7px' " +
        "fill='white' " +
        "width='%dpx' " +
        "height='%dpx'>\n",
        width, height)

    // Add SVG body skipping polygons with invalid corners
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            if ax, ay, err = corner(i + 1, j);     err { continue }
            if bx, by, err = corner(i,     j);     err { continue }
            if cx, cy, err = corner(i,     j + 1); err { continue }
            if dx, dy, err = corner(i + 1, j + 1); err { continue }
            fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }

    // Close SVG
    fmt.Fprintf(out, "</svg>\n")
}

// Main

func main() {
    genSVG()
}

// Processor 3.1 GHz Intel Core i7
// Memory: 16 GB 1867 MHz DDR3
// OS: OS X El Capitan v10.11.6
//
// BenchmarkGenSVG-4 (surface1) 33,811,283 ns/op (floating point output)
// BenchmarkGenSVG-4 (surface2) 12,047,030 ns/op (integer output)
