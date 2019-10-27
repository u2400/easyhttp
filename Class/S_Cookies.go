package Class

import "net/http"


type Cookies struct {
	Value []Cookie
}

func (cmap *Cookies) Init () {
	cmap.Value = []Cookie{}
}

func (cmap *Cookies) Insert (cookie Cookie) {
	cmap.Value = append(cmap.Value, cookie)
}

func (cmap * Cookies) Parser (Cookies []Cookie) {
	for _, value := range Cookies {
		if (value.Name == "" || value.Value == "" ) {
			panic("Must provide Cookie.Name and Cookie.value")
		}
		cmap.Insert(value)
	}
}

func (cmap * Cookies) AddCookieToReq (r *http.Request) {
	for _, value := range cmap.Value {
		r.AddCookie((*http.Cookie)(&value))
	}
}