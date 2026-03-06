package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		duration := time.Since(start)

		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		colorYellow := "\033[33m"
		colorRed := "\033[31m"
		colorCyan := "\033[36m"

		statusColor := colorGreen
		if wrapped.status >= 400 && wrapped.status < 500 {
			statusColor = colorYellow
		} else if wrapped.status >= 500 {
			statusColor = colorRed
		}

		log.Printf("%s[API]%s %s%-6s%s %s | %s%3d%s | %v",
			colorCyan, colorReset,
			colorCyan, r.Method, colorReset,
			r.URL.Path,
			statusColor, wrapped.status, colorReset,
			duration,
		)
	})
}
