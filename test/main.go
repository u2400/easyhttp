package main

import (
"fmt"
"github.com/u2400/easyhttp"
)

func main() {
	res := easyhttp.Get("http://baidu.com")
	fmt.Print(res)
}
