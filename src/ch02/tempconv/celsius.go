// Package tempconv performs Celsius and Fahrenheit temperature computations
package tempconv

import "fmt"

type Celsius    float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -237.15
    FreezingC     Celsius =    0.0
    BoilingC      Celsius =  100.0
)

func (c Celsius) String() string {
    return fmt.Sprintf("%g\u00b0C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g\u00b0F", f)
}

