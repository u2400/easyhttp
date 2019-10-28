package easyhttp

func Post(req interface{}) *HttpResponseStruct {
	res := forward(req, "POST")
	return res
}
