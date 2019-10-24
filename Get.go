package http

import (
	"./class"
	"./io_util"
	nativehttp "net/http"
)

func Get(method class.HttpRequestMethod) *class.HttpResponseMethod {
	client := &nativehttp.Client{}


	req, err := nativehttp.NewRequest("GET", method.Url, nil)
	if (err != nil) {
		return &class.HttpResponseMethod {
			Err: err,
		}
	}

	if (method.Header != nil) {
		method.MakeHeaderMapIntoReq(req)
	}

	cookie := nativehttp.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/"}
	req.AddCookie(&cookie)

	//send http request
	res,err := client.Do(req)
	defer res.Body.Close()
	if (err != nil) {
		return &class.HttpResponseMethod {
			Err: err,
		}
	}

	body_string, err := io_util.F_io_reader_to_string(&res.Body)

	if (err != nil) {
		return &class.HttpResponseMethod {
			Err: err,
		}
	}

	return &class.HttpResponseMethod {
		Url: method.Url,
		Body: body_string,
	}
}