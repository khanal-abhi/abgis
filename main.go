package main

import (
	"io"
	"strings"

	"github.com/khanal-abhi/abgis/config"
	amath "github.com/khanal-abhi/abgis/math"
	"github.com/khanal-abhi/jsonrpc2"
)

type abGISHandler struct {
	logger io.Reader
}

func (a abGISHandler) Handle(r jsonrpc2.Request) jsonrpc2.Response {
	methodChain := strings.Split(r.Method, "::")
	if len(methodChain) < 1 {
		return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
	}
	pkg := methodChain[0]
	switch pkg {
	case "distance":
		break
	case "math":
		return amath.Handle(r, methodChain[1:])
	case "notification":
		break
	default:
		break
	}
	return jsonrpc2.ErrorResponse(r.ID, 400, config.InvalidMethod, "")
}

func main() {

}
