package mware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("final "))
	})

	stack := Build(final, middleware("one"), middleware("two"), middleware("three"))
	rec := httptest.NewRecorder()
	stack.ServeHTTP(rec, nil)

	assert.Equal(t, strings.TrimSpace(rec.Body.String()), "one two three final threeDeferred twoDeferred oneDeferred")
}

func middleware(body string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer w.Write([]byte(body + "Deferred "))
			w.Write([]byte(body + " "))
			h.ServeHTTP(w, r)
		})
	}
}
