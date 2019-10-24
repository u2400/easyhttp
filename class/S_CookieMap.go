package class

import "net/http"

type CookieMap struct {
	Cookies map[int]http.Cookie
	num int
}

func (cmap *CookieMap) Insert (cookie http.Cookie) {
	cmap.Cookies[cmap.num] = cookie
	cmap.num += 1
}
