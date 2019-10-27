package Class

import "net/http"

type HttpRequestStruct struct {
	Url    string
	Header map[string]string
	Method string
	Cookie Cookies
	JsonBody map[string]interface{}
	RawBody string
	FilePath string
}

func (method *HttpRequestStruct) MakeHeaderMapIntoReq(req *http.Request) {
	for k,v := range method.Header {
		req.Header.Add(k,v)
	}
}
