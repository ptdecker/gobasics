// Package genconv defines types, constants, and conversions of units
package genconv

import "fmt"

type Celsius    float64
type Fahrenheit float64
type Kelvin     float64

const (
    AbsoluteZeroC Celsius = -237.15
    FreezingC     Celsius =    0.0
    BoilingC      Celsius =  100.0
)

func (c Celsius) String() string {
    return fmt.Sprintf("%.2f\u00b0C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%.2f\u00b0F", f)
}

func (k Kelvin) String() string {
    return fmt.Sprintf("%.2f K", k)
}

