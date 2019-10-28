package easyhttp

func Get(req interface{}) *HttpResponseStruct {
	res := forward(req, "GET")
	return res
}