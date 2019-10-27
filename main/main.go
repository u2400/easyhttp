package main

import (
	"fmt"
	"github.com/u2400/easyhttp"
)

func main() {
	//res := easyhttp.Get("http://47.96.232.4/show_server.php?a=123&b=321")
	res := easyhttp.Get(easyhttp.HttpRequestStruct{
		Url:"http://47.96.232.4/show_server.php?a=123&b=321",
	})
	fmt.Print(res)
}
