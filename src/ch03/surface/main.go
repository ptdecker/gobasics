// Surface computes an SVG rendering of a 3-D surface function.

package main

import (
    "fmt"
    "math"
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

var sin30 = math.Sin(angle)
var cos30 = math.Cos(angle)

// Definition of 3-D surface function to render
func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

// Return coordinates of the corner of a cell on a 2-D canvas
func corner(i, j int) (float64, float64) {
    
    // Find point (x, y) at corner of cell (i,j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute the surface height z using the surface function
    z := f(x, y)

    // Project (x, y, z) isometrically onto the 2-D SVG canvas (sx, sy)
    sx := width  / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return sx, sy
}

// Generate the SVG
func genSVG() {

    // Open SVG tag
    fmt.Printf("<svg " +
        "xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' " +
        "height='%d'>\n",
        width, height)

    // Add SVG body
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i + 1, j)
            bx, by := corner(i,     j)
            cx, cy := corner(i,     j + 1)
            dx, dy := corner(i + 1, j + 1)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }

    // Close SVG
    fmt.Println("</svg>")
}

// Main

func main() {
    genSVG()
}

