package distance

import (
	"errors"
	"math"

	amath "github.com/khanal-abhi/abgis/math"
)

// Constants related to the Earth
const (
	EarthRadius float64 = 6371000
)

// Point is an array that contains X, Y, ... Coordinates of the point
type Point []float64

// Euclidean returns the Euclidean distance between two Point types
func Euclidean(p1 Point, p2 Point, l int) (float64, error) {
	err := validateLength(p1, p2, l)
	if err != nil {
		return 0, err
	}
	diffSum := 0.0
	for i := 0; i < l; i++ {
		diff := p2[i] - p1[i]
		diffSum += diff * diff
	}
	sqrtDiffSum := math.Sqrt(diffSum)
	return sqrtDiffSum, nil
}

// Haversine returns the Haversine distance between two long lat Point Types in decimal degrees
func Haversine(p1 Point, p2 Point) (float64, error) {
	err := validateLength(p1, p2, 2)
	if err != nil {
		return 0, err
	}
	lat1, lng1 := amath.DegreesToRadians(p1[1]), amath.DegreesToRadians(p1[0])
	lat2, lng2 := amath.DegreesToRadians(p2[1]), amath.DegreesToRadians(p2[0])
	hDlat := (lat2 - lat1) / 2
	hDlng := (lng2 - lng1) / 2
	a := sinSquare(hDlat) + (math.Cos(lat1) * math.Cos(lat2) * sinSquare(hDlng))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := EarthRadius * c
	// leftTerm := sinSquare((lat2 - lat1) / 2)
	// rightTerm := math.Cos(lng1) * math.Cos(lng2) * sinSquare((lng2-lng1)/2)
	// rootTerm := math.Sqrt(leftTerm + rightTerm)
	// result := 2 * EarthRadius * math.Asin(rootTerm)
	return d, nil
}

func sinSquare(v float64) float64 {
	sin := math.Sin(v)
	return sin * sin
}

func validateLength(p1 Point, p2 Point, l int) error {
	len1 := len(p1)
	len2 := len(p2)
	if len1 < l || len2 < l {
		return errors.New("not enough co-ordinates in one of the points")
	}
	return nil
}
