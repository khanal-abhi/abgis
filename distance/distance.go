package distance

import (
	"encoding/json"
	"errors"
	"math"

	"github.com/khanal-abhi/abgis/config"
	amath "github.com/khanal-abhi/abgis/math"
	"github.com/khanal-abhi/jsonrpc2"
)

// Constants related to the Earth
const (
	EarthRadius           float64 = 6371000
	NotEnoughCoordinates  string  = "NOT_ENOUGH_COORDINATES"
	InvalidNumberOfPoints string  = "INVALID_NUMBER_OF_POINTS"
)

// Handle handles all the jsonrpc2 traffic to distance
func Handle(r jsonrpc2.Request, methodChain []string) jsonrpc2.Response {
	if len(methodChain) < 1 {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
	}
	switch methodChain[0] {
	case "Haversine":
		return handleHaversine(r)
	case "Euclidean":
		return handleEuclidean(r)
	default:
		break
	}
	return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
}

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
		return errors.New(NotEnoughCoordinates)
	}
	return nil
}

func extractPoints(params string) ([]Point, error) {
	pts := make([]Point, 0)
	ps := []byte(params)
	err := json.Unmarshal(ps, &pts)
	if err != nil {
		return nil, err
	}
	if len(pts) != 2 {
		return pts, errors.New(InvalidNumberOfPoints)
	}
	return pts, nil
}

func handleHaversine(r jsonrpc2.Request) jsonrpc2.Response {
	pts, err := extractPoints(r.Params)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}

	h, err := Haversine(pts[0], pts[1])
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}

	b, err := json.Marshal([]float64{h})
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	return jsonrpc2.NewResponse(r.ID, string(b), jsonrpc2.Error{})
}

func handleEuclidean(r jsonrpc2.Request) jsonrpc2.Response {
	pts, err := extractPoints(r.Params)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	l1, l2 := float64(len(pts[0])), float64(len(pts[1]))
	l := int(math.Min(l1, l2))
	e, err := Euclidean(pts[0], pts[1], l)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}

	b, err := json.Marshal([]float64{e})
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	return jsonrpc2.NewResponse(r.ID, string(b), jsonrpc2.Error{})
}
