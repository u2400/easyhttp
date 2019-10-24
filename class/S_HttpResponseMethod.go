package class

type HttpResponseMethod struct {
	StatusCode int8
	Url string
	Headers map[string]string
	Body string
	Err error
	encode string
	Text string
	Json map[string]interface{}
}


