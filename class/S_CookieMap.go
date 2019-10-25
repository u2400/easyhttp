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

func (cmap * CookieMap) Parser (Cookies []http.Cookie) {
	for _, value := range Cookies {
		if (value.Name == "" || value.Value == "" ) {
			panic("Must provide Cookie.Name and Cookie.value")
		}
		cmap.Insert(value)
	}
}

func (cmap * CookieMap) AddCookieToReq (r *http.Request) {
	for _, value := range cmap.Cookies {
		r.AddCookie(&value)
	}
}