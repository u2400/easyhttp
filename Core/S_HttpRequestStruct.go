package Core

import (
	"net/http"
)

type HttpRequestStruct struct {
	Url      string
	Header   map[string]string
	Method   string
	Cookies  Cookies
	JsonBody map[string]interface{}
	RawBody  string
	FilePath string
}

func MakeHeaderMapIntoReq(this *HttpRequestStruct, req *http.Request) {
	for k,v := range this.Header {
		req.Header.Add(k,v)
	}
}
