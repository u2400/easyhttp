package Core

import (
	nativehttp "net/http"
	"strings"
)

func Request(UserRequest []byte, HookMap map[HookName]func(HookObject HookStruct)) []byte {
	var HttpRequestObject HttpRequestStruct = toCoreRequest(UserRequest, HookMap)
	var HttpResponseObject HttpResponseStruct = HttpResponseStruct{
		Url:HttpRequestObject.Url,
	}

	client := &nativehttp.Client{}

	req, err := nativehttp.NewRequest(HttpRequestObject.Method, HttpRequestObject.Url, nil)
	if (err != nil) {
		panic(err)
	}

	if (HttpRequestObject.Header != nil) {
		MakeHeaderMapIntoReq(&HttpRequestObject, req)
	}

	//set Cookie
	HttpRequestObject.Cookies.AddCookieToReq(req)

	//send http request
	res,err := client.Do(req)
	defer res.Body.Close()
	if (err != nil) {
		panic(err)
	}

	//get io reader body to string
	body_string, err := ioReaderToString(&res.Body)
	HttpResponseObject.Body = body_string

	if (err != nil) {
		panic(err)
	}

	Headers := map[string]string{}
	for k, v := range res.Header {
		Headers[k] = strings.Join(v,"")
	}
	HttpResponseObject.Headers = Headers

	return toUserResponse(HttpResponseObject)
}