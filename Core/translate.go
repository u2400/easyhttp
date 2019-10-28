package Core

import (
	"encoding/json"
)

func toCoreRequest(_r []byte) HttpRequestStruct {
	var HttpRequestObject HttpRequestStruct = HttpRequestStruct{}

	//make json to object
	err := json.Unmarshal(_r, &HttpRequestObject)
	if (err != nil) {
		panic(err)
	}

	return HttpRequestObject
}

func toUserResponse(_r HttpResponseStruct) []byte {
	s, err := json.Marshal(_r)
	if (err != nil) {
		panic(err)
	}

	return s
}