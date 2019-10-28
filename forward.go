package easyhttp

import (
	"encoding/json"
	"github.com/u2400/easyhttp/Core"
	"reflect"
)

func forward(_req interface{},method string) *HttpResponseStruct {
	var req *HttpRequestStruct
	var res *HttpResponseStruct
	switch reflect.TypeOf(_req) {

	//support stronger http request
	case reflect.TypeOf(HttpRequestStruct{}):
		tmp := ((_req).(HttpRequestStruct))
		req = &tmp

	//support quick http request
	case reflect.TypeOf(""):
		var url string = _req.(string)
		req = &HttpRequestStruct{
			Url:url,
			Method: method,
		}

	default:
		error := "\n[from forward.go->(func forward)] unexpect arg type: expect HttpRequestStruct{} or string\n"
		error += "[from forward.go->(func forward)] 非预期的参数类型: 需要一个HttpRequestStruct{}或string类型参数"
		panic(error)
	}

	JsonReq, err := json.Marshal(req)
	if (err != nil) {
		panic(err)
	}

	JsonRes := Core.Request(JsonReq)

	err = json.Unmarshal(JsonRes, &res)
	if (err != nil) {
		panic(err)
	}

	return res
}