package math

import (
	"encoding/json"
	"math"

	"github.com/khanal-abhi/abgis/config"
	"github.com/khanal-abhi/jsonrpc2"
)

// Handle handles all the jsonrpc2 traffic to math
func Handle(r jsonrpc2.Request, methodChain []string) jsonrpc2.Response {
	if len(methodChain) < 1 {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
	}
	switch methodChain[0] {
	case "FloatCompare":
		return handleFloatCompare(r)
	case "DegreesToRadians":
		return handleDegreesToRadians(r)
	case "RadiansToDegrees":
		return handleRadiansToDegrees(r)
	default:
		break
	}
	return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
}

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

func handleFloatCompare(r jsonrpc2.Request) jsonrpc2.Response {
	fArray := make([]float64, 3)
	err := json.Unmarshal([]byte(r.Params), &fArray)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	fc := FloatCompare(fArray[0], fArray[1], fArray[2])
	b, err := json.Marshal([]int{fc})
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	return jsonrpc2.NewResponse(r.ID, string(b), jsonrpc2.Error{})
}

func handleDegreesToRadians(r jsonrpc2.Request) jsonrpc2.Response {
	fArray := make([]float64, 1)
	err := json.Unmarshal([]byte(r.Params), &fArray)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	d2r := DegreesToRadians(fArray[0])
	b, err := json.Marshal([]float64{d2r})
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	return jsonrpc2.NewResponse(r.ID, string(b), jsonrpc2.Error{})
}

func handleRadiansToDegrees(r jsonrpc2.Request) jsonrpc2.Response {
	fArray := make([]float64, 1)
	err := json.Unmarshal([]byte(r.Params), &fArray)
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	r2d := RadiansToDegrees(fArray[0])
	b, err := json.Marshal([]float64{r2d})
	if err != nil {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidParameters, "")
	}
	return jsonrpc2.NewResponse(r.ID, string(b), jsonrpc2.Error{})
}
