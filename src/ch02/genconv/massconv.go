// Package genconv defines types, constants, and conversions of units
package genconv

// LbToKg converts a pounds to kilograms
func LbToKg(lb Pound) Kilogram {
    return Kilogram(lb * 0.454)
}

// KgToLb converts a kilograms to pounds
func KgToLb(kg Kilogram) Pound {
    return Pound(kg * 2.205)
}

