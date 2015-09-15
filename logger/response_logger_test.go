package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Write(t *testing.T) {
	rec := httptest.NewRecorder()
	rl := NewResponseLogger(rec)

	rl.Write([]byte("Lube job while you wait, Govna?"))
	rl.Write([]byte(" Don't touch me."))

	assert.Equal(t, rl.Status(), http.StatusOK)
	assert.Equal(t, rec.Code, rl.Status())
	assert.Equal(t, rec.Body.String(), "Lube job while you wait, Govna? Don't touch me.")
	assert.Equal(t, rl.Size(), 47)
}

func Test_WriteHeader(t *testing.T) {
	rec := httptest.NewRecorder()
	rl := NewResponseLogger(rec)

	rl.WriteHeader(http.StatusNoContent)

	assert.Equal(t, rl.Status(), http.StatusNoContent)
	assert.Equal(t, rec.Code, rl.Status())
}

func Test_Flush(t *testing.T) {
	rec := httptest.NewRecorder()
	rl := NewResponseLogger(rec)
	rl.Flush()

	assert.True(t, rec.Flushed)
}
