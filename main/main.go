package main

import (
	"fmt"
	"github.com/u2400/easyhttp"
	"github.com/u2400/easyhttp/Class"
)

func main() {
	cookie := Class.Cookies{}
	cookie.Init()
	cookie.Insert(Class.Cookie{
		Name:"test",
		Value:"test",
	})
	res := easyhttp.Get("http://47.96.232.4/show_server.php?a=123&b=321")
	fmt.Print(res)
}
