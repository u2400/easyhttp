package easyhttp

import (
	"github.com/u2400/easyhttp/Class"
	nativehttp "net/http"
	"strings"
)

func Request(HttpRequestMethod Class.HttpRequestMethod) *Class.HttpResponseMethod {
	client := &nativehttp.Client{}

	req, err := nativehttp.NewRequest(HttpRequestMethod.Method, HttpRequestMethod.Url, nil)
	if (err != nil) {
		return &Class.HttpResponseMethod {
			Err: err,
		}
	}

	if (HttpRequestMethod.Header != nil) {
		HttpRequestMethod.MakeHeaderMapIntoReq(req)
	}

	HttpRequestMethod.Cookie.AddCookieToReq(req)

	//send http request
	res,err := client.Do(req)
	defer res.Body.Close()
	if (err != nil) {
		return &Class.HttpResponseMethod {
			Err: err,
		}
	}

	body_string, err := Class.IoReaderToString(&res.Body)

	if (err != nil) {
		return &Class.HttpResponseMethod {
			Err: err,
		}
	}

	Headers := map[string]string{}
	for k, v := range res.Header {
		Headers[k] = strings.Join(v,"")
	}

	return &Class.HttpResponseMethod {
		Url: HttpRequestMethod.Url,
		StatusCode:res.StatusCode,
		Body: body_string,
		Headers: Headers,
	}
}