package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("START")

	r, e := ReadBox()
	if e.IsError() {
		fmt.Println(e.Error)
	}

	js, _ := json.Marshal(r)
	fmt.Println(string(js))

	r2, e := ReadRun()
	if e.IsError() {
		fmt.Println(e.Error)
	}

	js, _ = json.Marshal(r2)
	fmt.Println(string(js))

	r3, e := RunFunc(RunBoxRequest{
		BoxId:        1,
		FunctionName: "fib",
		Params: []FunctionParameter{{
			Value: 1,
			Ptype: INT32,
		}},
	})
	if e.IsError() {
		fmt.Println(e.Error)
	}

	js, _ = json.Marshal(r3)
	fmt.Println(string(js))

	var server Server
	server.Start()

	fmt.Println("END")
}

func ReadRun() ([]Box, Info) {
	var respStruct []Box

	resp, err := http.Get(BASE_PROTOCOL + "://" + BASE_ADDR + ":" + BASE_PORT + "/" + BASE_WASM + "/" + "run")
	if err != nil {
		return respStruct, InfoError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respStruct, InfoError(err)
	}

	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return respStruct, InfoError(err)
	}

	return respStruct, InfoOk()
}

func ReadBox() ([]BoxSimple, Info) {
	var respStruct []BoxSimple

	resp, err := http.Get(BASE_PROTOCOL + "://" + BASE_ADDR + ":" + BASE_PORT + "/" + BASE_WASM + "/" + "run")
	if err != nil {
		return respStruct, InfoError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respStruct, InfoError(err)
	}

	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return respStruct, InfoError(err)
	}

	return respStruct, InfoOk()
}

func RunFunc(runReq RunBoxRequest) (RunBoxResponse, Info) {
	var respStruct RunBoxResponse

	json_data, err := json.Marshal(runReq)

	// RunBoxRequest{
	// 	BoxId:        1,
	// 	FunctionName: "fib",
	// 	Params: []FunctionParameter{{
	// 		Value: 1,
	// 		Ptype: INT32,
	// 	}},
	// }

	fmt.Println(string(json_data))

	resp, err := http.Post(BASE_PROTOCOL+"://"+BASE_ADDR+":"+BASE_PORT+"/"+BASE_WASM+"/"+"run", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return respStruct, InfoOk()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respStruct, InfoOk()
	}

	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return respStruct, InfoError(err)
	}

	return respStruct, InfoOk()
}
