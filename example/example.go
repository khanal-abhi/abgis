package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	jsonrpc2 "github.com/khanal-abhi/jsonrpc2"
)

type handlr struct{}

func (h handlr) Handle(r jsonrpc2.Request) jsonrpc2.Response {
	return jsonrpc2.NewResponse(r.ID, "", jsonrpc2.Error{})
}

func main() {
	s := jsonrpc2.Server{}
	handlersMap := make(map[string]jsonrpc2.IHandler)
	handlersMap["a"] = handlr{}
	wg := sync.WaitGroup{}
	sDone := make(chan bool)
	wg.Add(1)
	go server(s, handlersMap, &wg, sDone)
	wg.Wait()
	wg.Add(1)
	go client(&wg)
	wg.Wait()
	<-sDone
}

func server(s jsonrpc2.Server, handlersMap map[string]jsonrpc2.IHandler, wg *sync.WaitGroup, sDone chan<- bool) {
	wg.Done()
	err := s.Serve(":8080", handlersMap)
	if err != nil {
		fmt.Println(err)
	}
	sDone <- true

}

func client(wg *sync.WaitGroup) {
	c, err := net.Dial("tcp", "localhost:8080")
	jsonEncoder := json.NewEncoder(c)
	jsonDecoder := json.NewDecoder(c)
	if err != nil {
		fmt.Println(err)
	} else {
		req := jsonrpc2.Request{
			ID:      101,
			Method:  "a",
			JSONRpc: jsonrpc2.JSONRPCVersion,
			Params:  "",
		}
		err = jsonEncoder.Encode(req)
		if err != nil {
			fmt.Println(err)
		} else {
			res := jsonrpc2.Response{}
			err = jsonDecoder.Decode(&res)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			_ = c.Close()
		}
	}
	wg.Done()
}
