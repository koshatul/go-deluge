package deluge

import (
	"net/http"
	"net/url"
)

// CookieJar is a simple and *very* insecure cookiejar that returns all cookies for all requests.
type CookieJar struct {
	cookies []*http.Cookie
}

// NewCookieJar returns a new `CookieJar` object.
func NewCookieJar() *CookieJar {
	return &CookieJar{
		cookies: []*http.Cookie{},
	}
}

// Cookies implements the Cookies method of the http.CookieJar interface.
func (c *CookieJar) Cookies(u *url.URL) (cookies []*http.Cookie) {
	return c.cookies
}

// SetCookies implements the SetCookies method of the http.CookieJar interface.
func (c *CookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.cookies = append(c.cookies, cookies...)
}
