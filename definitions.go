package main

var BASE_PROTOCOL = "http"
var BASE_ADDR = "10.0.1.209"
var BASE_PORT = "10001"
var BASE_HARDWARE = "hardware"
var BASE_WASM = "wasm"
var BASE_RUN = "run"
var BASE_BOX = "box"

type RunBoxRequest struct {
	BoxId        int
	FunctionName string
	Params       []FunctionParameter
}

type RunBoxResponse struct {
	Response  []interface{}
	Error     int
	ErrorDesc string
}

type Box struct {
	Id            int
	Name          string
	FunctionCount int
	ModuleCount   int
	Functions     []FunctionDefinition
	ErrorId       bool
}

type BoxSimple struct {
	Id      int
	Name    string
	ErrorId bool
}

type FunctionDefinition struct {
	Name     string
	Index    int
	ArgCount int
	RetCount int
	Types    []FunctionParameterType
}

type FunctionParameter struct {
	Value interface{}
	Ptype FunctionParameterType
}

type FunctionParameterType int

// c_m3Type_none   = 0,
// c_m3Type_i32    = 1,
// c_m3Type_i64    = 2,
// c_m3Type_f32    = 3,
// c_m3Type_f64    = 4,

const (
	INT32 FunctionParameterType = 1
	INT64 FunctionParameterType = 2
	FLT32 FunctionParameterType = 3
	FLT64 FunctionParameterType = 4
	BOOLI FunctionParameterType = 5 // moj tip
	STPTR FunctionParameterType = 6 // moj tip
)

type Info struct {
	Message string
	Error   bool
}

func InfoOk() Info {
	return Info{
		Message: "",
		Error:   false,
	}
}

func InfoError(err error) Info {
	return Info{
		Message: err.Error(),
		Error:   true,
	}
}

func (i Info) IsError() bool {
	return i.Error
}

// [
//    {
//       "Id":1,
//       "Name":"Fib32",
//       "ErrorId":false,
//       "FunctionCount":1,
//       "ModuleCount":1,
//       "Functions":[
//          [
//             {
//                "Index":0,
//                "Name":"fib",
//                "ArgCount":1,
//                "RetCount":1,
//                "Types":[
//                   1,
//                   1
//                ]
//             }
//          ]
//       ]
//    },
//    {
//       "Id":2,
//       "Name":"Fib32_V2",
//       "ErrorId":false,
//       "FunctionCount":1,
//       "ModuleCount":1,
//       "Functions":[
//          [
//             {
//                "Index":0,
//                "Name":"fib",
//                "ArgCount":1,
//                "RetCount":1,
//                "Types":[
//                   1,
//                   1
//                ]
//             }
//          ]
//       ]
//    }
// ]
