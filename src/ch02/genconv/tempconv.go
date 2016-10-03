// Package genconv defines types, constants, and conversions of units
package genconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin {
    return Kelvin(c + 273.15)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin {
    return Kelvin((f + 459.67) * 5 / 9)
}

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius {
    return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
    return Fahrenheit(k * 9 / 5 - 459.67)
}

