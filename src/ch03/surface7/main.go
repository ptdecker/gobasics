// Surface7 computes a colored SVG rendering of a 3-D surface function.  It also includes guard code to
// ignore any polygons whose corner values are infinte or NaN.  Colors run from red at the peaks to
// blue in the valley.  It also includes a built-in web server mode.

// NOTE: This solution approach utilizes concepts from Chapter 4.  I could not come up with a clean
//       way of solving it that did not reference concepts not yet introduced in Chapther 3.  However,
//       they are touched on in Chapter 1 to some extent.  Specifically, this solution utilizes
//       structures and arrays.

// c.f. https://developer.mozilla.org/en-US/docs/Web/SVG
// c.f. https://jwatt.org/svg/authoring/

package main

import (
    "fmt"
    "io"
    "log"
    "math"
    "net/http"
    "os"
)

const (
    width   = 600                  // canvas width in pixels
    height  = 320                  // canvas height in pixels
    cells   = 100                  // number of grid cells
    xyrange = 30.0                 // axis ranges (-xyrange..+xyrange)
    xyscale = width / 2 / xyrange  // pixels per x or y unit
    zscale  = height * 0.4         // pixels per z unit
    angle   = math.Pi / 6          // angle of x and y axis (30 deg)
)

// Structure to store polygon coordinates
type Poly struct {
    ax, ay  float64
    bx, by  float64
    cx, cy  float64
    dx, dy  float64
    zavg    float64
} 

var out io.Writer = os.Stdout   // modified during testing
var sin30 = math.Sin(angle)
var cos30 = math.Cos(angle)
var polygons [cells][cells]Poly // Array of polygons

// Definition of 3-D surface function to render
func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

// Return coordinates of the corner of a cell on a 2-D canvas along with its z value
func corner(i, j int) (float64, float64, float64, bool) {
    
    // Find point (x, y) at corner of cell (i,j)
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    // Compute the surface height z using the surface function
    z := f(x, y)

    // Project (x, y, z) isometrically onto the 2-D SVG canvas (sx, sy)
    sx := width  / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return sx, sy, z, (math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)) 
}

// Calculate the cells
func genCells(cells int) (float64, float64) {

    var z1, z2, z3, z4 float64
    var err bool
    zmin := 0.0
    zmax := 0.0

    // Add SVG body skipping polygons with invalid corners
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            if polygons[i][j].ax, polygons[i][j].ay, z1, err = corner(i + 1, j);     err { continue }
            if polygons[i][j].bx, polygons[i][j].by, z2, err = corner(i,     j);     err { continue }
            if polygons[i][j].cx, polygons[i][j].cy, z3, err = corner(i,     j + 1); err { continue }
            if polygons[i][j].dx, polygons[i][j].dy, z4, err = corner(i + 1, j + 1); err { continue }
            zavg := (z1 + z2 + z3 + z4) / 4.0
            polygons[i][j].zavg = zavg
            if !math.IsNaN(zavg) {
                zmin = math.Min(zmin, zavg)
                zmax = math.Max(zmax, zavg)
            }

        }
    }

    return zmin, zmax

 }


// Generate the SVG
func genSVG(out io.Writer) {

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

    // Print polygons
    zmin, zmax := genCells(cells)

    // Add SVG body skipping polygons with invalid corners
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            red   := int(math.Floor(255 * ((polygons[i][j].zavg - zmin) / (zmax - zmin))))
            green := 0
            blue  := 255 - red
            fmt.Fprintf(out, "<polygon fill='rgb(%d,%d,%d)' stroke='rgb(%d,%d,%d)' points='%g,%g %g,%g %g,%g %g,%g' />\n",
                red, green, blue,
                red, green, blue,
                polygons[i][j].ax, polygons[i][j].ay,
                polygons[i][j].bx, polygons[i][j].by,
                polygons[i][j].cx, polygons[i][j].cy,
                polygons[i][j].dx, polygons[i][j].dy)
        }
    }

    // Close SVG
    fmt.Fprintf(out, "</svg>\n")
}

// Launch a basic web server
func webServer() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        genSVG(w)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// Main

func main() {

    // If called with "web" as a parameter, launch web server
    // otherwise output an SVG

    if len(os.Args) > 1 && os.Args[1] == "web" {
        webServer()
    } else {
        genSVG(out)
    }
}

// Processor 3.1 GHz Intel Core i7
// Memory: 16 GB 1867 MHz DDR3
// OS: OS X El Capitan v10.11.6
//
// BenchmarkGenSVG-4 (surface1) 33,811,283 ns/op (floating point output)
// BenchmarkGenSVG-4 (surface2) 12,047,030 ns/op (integer output)
