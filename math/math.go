package math

import "math"

// FloatCompare compares two float numbers with a threshold
func FloatCompare(f1 float64, f2 float64, threshold float64) int {
	diff := f1 - f2
	diffSign := 1
	if diff < 0 {
		diffSign = -1
	}
	aDiff := math.Abs(diff)
	if aDiff <= threshold {
		return 0
	}
	return diffSign
}

// DegreesToRadians converts decimal degrees to radians
func DegreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// RadiansToDegrees converts radians to decimal degrees
func RadiansToDegrees(r float64) float64 {
	return r * 180 / math.Pi
}
