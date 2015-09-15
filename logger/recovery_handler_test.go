package logger

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Recover(t *testing.T) {
	buff := bytes.NewBufferString("")
	logger := log.New(buff, "", 0)
	rec := httptest.NewRecorder()

	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("panic!")
	})

	Recover(logger)(final).ServeHTTP(rec, nil)

	assert.Equal(t, rec.Code, http.StatusInternalServerError)
	assert.Equal(t, buff.String(), "panic!\n")
}
