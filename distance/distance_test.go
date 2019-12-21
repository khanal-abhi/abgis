package distance

import (
	"testing"

	"github.com/khanal-abhi/jsonrpc2"

	amath "github.com/khanal-abhi/abgis/math"
)

func TestHandle(t *testing.T) {
	r := jsonrpc2.Request{
		ID:      1,
		Method:  "",
		Params:  "[[1, 0], [0, 1]]",
		JSONRpc: jsonrpc2.JSONRPCVersion,
	}
	res := Handle(r, []string{"Euclidean"})
	if res.Error.Code != 0 {
		t.Fail()
	}
}

func TestEuclidean(t *testing.T) {
	th := 0.001
	pa1 := []float64{1, 2}
	pa2 := []float64{3, -4}
	pad := 6.324555

	da, err := Euclidean(pa1, pa2, 2)
	if err != nil {
		t.Fail()
	}
	cmp := amath.FloatCompare(da, pad, th)
	if cmp != 0 {
		t.Fail()
	}

	pa3 := []float64{1, 2, 3}
	da, err = Euclidean(pa1, pa3, 3)
	if err == nil {
		t.Fail()
	}
}

func TestHaversine(t *testing.T) {
	th := 1000.0
	pa1 := []float64{-21.8174, 64.1265}
	pa2 := []float64{-74.0060, 40.7128}
	pad := 4208000.0

	da, err := Haversine(pa1, pa2)
	if err != nil {
		t.Fail()
	}
	cmp := amath.FloatCompare(da, pad, th)
	if cmp != 0 {
		t.Fail()
	}

	pa3 := []float64{1, 2, 3}
	da, err = Euclidean(pa1, pa3, 3)
	if err == nil {
		t.Fail()
	}
}
