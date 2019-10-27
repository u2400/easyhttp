package easyhttp

import (
"github.com/u2400/easyhttp/Class"
"reflect"
)

func Get(method interface{}) *Class.HttpResponseMethod {
	var res *Class.HttpResponseMethod
	switch reflect.TypeOf(method) {

	case reflect.TypeOf(Class.HttpRequestMethod{}):
		res = Request(method.(Class.HttpRequestMethod))
	case reflect.TypeOf(""):
		res = Request(Class.HttpRequestMethod{
			Url:method.(string),
			Method: "GET",
		})
	}

	return res
}