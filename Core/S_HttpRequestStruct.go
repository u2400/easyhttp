package Core

import (
	"net/http"
)

type HookName string
type HttpRequestStruct struct {
	Url      string
	Header   map[string]string
	Method   string
	Cookies  Cookies
	JsonBody map[string]interface{}
	RawBody  string
	FilePath string
	HookMap map[HookName]func(HookObject HookStruct)
}

func MakeHeaderMapIntoReq(this *HttpRequestStruct, req *http.Request) {
	for k,v := range this.Header {
		req.Header.Add(k,v)
	}
}

//func GetAttributes( _r UserRequestInterface) HttpRequestStruct{
//	req := HttpRequestStruct{}
//	v := reflect.ValueOf(_r)
//	_req := v.
//	count := v.NumField()
//
//	for i := 0 ; i < count ; i++ {
//		fmt.Println(v.Elem().Field(i))
//	}
//	fmt.Print(req)
//	return req
//}
