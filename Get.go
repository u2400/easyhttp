package easyhttp

import (
	"github.com/u2400/easyhttp/Class"
	"reflect"
)

func Get(method interface{}) *Class.HttpResponseStruct {
	var res *Class.HttpResponseStruct
	switch reflect.TypeOf(method) {

	case reflect.TypeOf(Class.HttpRequestStruct{}):
		var method Class.HttpRequestStruct = method.(Class.HttpRequestStruct)
		method.Method = "GET"
		res = Request(method)
	case reflect.TypeOf(""):
		var method string = method.(string)
		res = Request(Class.HttpRequestStruct{
			Url:method,
			Method: "GET",
		})
	}

	return res
}