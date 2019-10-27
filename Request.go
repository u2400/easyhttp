package easyhttp

import (
	"github.com/u2400/easyhttp/Class"
	nativehttp "net/http"
	"strings"
)

func Request(HttpRequestStruct Class.HttpRequestStruct) *Class.HttpResponseStruct {
	client := &nativehttp.Client{}

	req, err := nativehttp.NewRequest(HttpRequestStruct.Method, HttpRequestStruct.Url, nil)
	if (err != nil) {
		return &Class.HttpResponseStruct {
			Err: err,
		}
	}

	if (HttpRequestStruct.Header != nil) {
		HttpRequestStruct.MakeHeaderMapIntoReq(req)
	}

	HttpRequestStruct.Cookie.AddCookieToReq(req)

	//send http request
	res,err := client.Do(req)
	defer res.Body.Close()
	if (err != nil) {
		return &Class.HttpResponseStruct {
			Err: err,
		}
	}

	body_string, err := Class.IoReaderToString(&res.Body)

	if (err != nil) {
		return &Class.HttpResponseStruct {
			Err: err,
		}
	}

	Headers := map[string]string{}
	for k, v := range res.Header {
		Headers[k] = strings.Join(v,"")
	}

	return &Class.HttpResponseStruct {
		Url: HttpRequestStruct.Url,
		StatusCode:res.StatusCode,
		Body: body_string,
		Headers: Headers,
	}
}