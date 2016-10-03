// Package genconv defines types, constants, and conversions of units
package genconv

import "fmt"

type Pound     float64
type Kilogram  float64

func (lb Pound) String() string {
    return fmt.Sprintf("%.2f lb", lb)
}

func (kg Kilogram) String() string {
    return fmt.Sprintf("%.2f kg", kg)
}

