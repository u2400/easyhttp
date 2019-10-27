package Class

type HttpResponseMethod struct {
	StatusCode int
	Url string
	Headers map[string]string
	Body string
	Err error
	encode string
	Text string
	Json map[string]interface{}
}


