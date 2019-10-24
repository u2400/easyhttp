package class

import "net/http"

type HttpRequestMethod struct {
	Url    string
	Header map[string]string
	Cookie CookieMap
	JsonBody map[string]interface{}
	RawBody string
	FilePath string
}

func (method *HttpRequestMethod) MakeHeaderMapIntoReq(req *http.Request) {
	for k,v := range method.Header {
		req.Header.Add(k,v)
	}
}
