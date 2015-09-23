package mware

import "net/http"

// Returns a handler composed of a slice of Middleware with final http.Handler as the last handler in the stack.
// The handler returned executes Middleware in the same order as the slice.
//
// For clarity, a verbose implementation could look like:
//
// func Build(final http.Handler, middlewareSlice ...Middleware) http.Handler {
//	composedHandler = final
// 	for i := len(middlewareSlice) - 1; i >= 0; i-- {
//		middleware = middlewareSlice[i]
// 		composedHandler = middleWare(composedHandler)
// 	}
// 	return composedHandler
// }
//
//
func Build(final http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		final = m[i](final)
	}
	return final
}

// Delegates to Build(final http.Handler, m ...Middleware)
// Cuts down on boilerplate. Prevents this common wrapping pattern: http.HandlerFunc(func (w http.ResponseWriter, r *http.Request))
// With BuildFunc, you can pass any final that conforms to the signature func (w http.ResponseWriter, r *http.Request)
//
func BuildFunc(final http.HandlerFunc, m ...Middleware) http.Handler {
	return Build(final, m...)
}
