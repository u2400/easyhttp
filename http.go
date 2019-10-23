package http

import (
	ioutil "./io_util"
	nativehttp "net/http"
)

type S_Http_request_method struct {
	Url    string
	Header map[string]string
}

type S_Http_response_method struct {
	StatusCode int8
	Url string
	Header map[string]string
	Body string
	Err error
}

type cookie nativehttp.Cookie
type CookieMap struct {
	Cookies map[int]nativehttp.Cookie
	num int
}

func (cmap *CookieMap)insert (cookie nativehttp.Cookie){
	cmap.Cookies[cmap.num] = cookie
	cmap.num += 1
}

func make_header_map_into_req(method *S_Http_request_method, req *nativehttp.Request) {
	for k,v := range method.Header {
		req.Header.Add(k,v)
	}
}

func Get(method S_Http_request_method) *S_Http_response_method {
	client := &nativehttp.Client{}

	req, err := nativehttp.NewRequest("GET", method.Url, nil)
	if (err != nil) {
		return &S_Http_response_method {
			Err: err,
		}
	}

	make_header_map_into_req(&method, req)
	cookie := nativehttp.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/"}
	req.AddCookie(&cookie)

	//send http request
	res,err := client.Do(req)
	defer res.Body.Close()
	if (err != nil) {
		return &S_Http_response_method {
			Err: err,
		}
	}

	body_string, err := ioutil.F_io_reader_to_string(&res.Body)
	if (err != nil) {
		return &S_Http_response_method {
			Err: err,
		}
	}

	return &S_Http_response_method {
		Url: method.Url,
		Body: body_string,
	}
}

//func Post(method S_Http_request_method) *S_Http_response_method {
//	return
//}

