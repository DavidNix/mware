package logger

import "net/http"

// ResponseLogger is a wrapper around http.ResponseWriter that captures extra information about
// the response.
type ResponseLogger interface {
	http.ResponseWriter
	http.Flusher
	// Status returns the status code of the response or 0 if the response has not been written.
	Status() int
	// Size returns the size in bytes of the response body.
	Size() int
}

// NewResponseLogger creates a ResponseLogger that wraps an http.ResponseWriter
func NewResponseLogger(w http.ResponseWriter) ResponseLogger {
	return &responseLogger{w, 0, 0}
}

type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (rl *responseLogger) Header() http.Header {
	return rl.w.Header()
}

func (rl *responseLogger) Write(b []byte) (int, error) {
	if rl.status == 0 {
		rl.WriteHeader(http.StatusOK)
	}
	size, err := rl.w.Write(b)
	rl.size += size
	return size, err
}

func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status
	rl.w.WriteHeader(status)
}

func (rl *responseLogger) Status() int {
	return rl.status
}

func (rl *responseLogger) Size() int {
	return rl.size
}

func (rl *responseLogger) Flush() {
	flusher, ok := rl.w.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}
