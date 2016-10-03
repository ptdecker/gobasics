// Package genconv defines types, constants, and conversions of units
package genconv

import "fmt"

type Foot   float64
type Meter  float64

func (ft Foot) String() string {
    return fmt.Sprintf("%.2f ft", ft)
}

func (m Meter) String() string {
    return fmt.Sprintf("%.2f m", m)
}

