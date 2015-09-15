package mware

import "net/http"

func Build(final http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		final = m[i](final)
	}
	return final
}

func BuildFunc(final http.HandlerFunc, m ...Middleware) http.Handler {
	return Build(final, m...)
}
