// Package genconv defines types, constants, and conversions of units
package genconv

// FtToM converts feet to meters
func FtToM(ft Foot) Meter {
    return Meter(ft * 0.305)
}

// MToFt converts a meters to feet
func MToFt(m Meter) Foot {
    return Foot(m * 3.281)
}

