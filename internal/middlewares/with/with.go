package with

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ufukty/bs/internal/colors"
)

type responseWriter struct {
	http.ResponseWriter
	StartTime  time.Time
	StatusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

var _ http.ResponseWriter = (*responseWriter)(nil)

func summarizeRequest(r *http.Request) string {
	return fmt.Sprintf("%s %s %s (%s, %s)",
		colors.Green(r.Method),
		colors.Yellow(r.URL.Path),
		colors.Red(r.Proto),
		colors.Blue(r.Host),
		colors.Magenta(r.RemoteAddr),
	)
}

func summarizeResponse(w *responseWriter, t time.Time) string {
	return fmt.Sprintf("%s %s %s bytes",
		colors.Magenta(w.StatusCode),
		colors.Green(time.Since(t)),
		colors.Cyan(w.Header().Get("Content-Length")),
	)
}

func Logging(h http.Handler) http.Handler {
	count := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		rw := &responseWriter{ResponseWriter: w, StartTime: t}
		fmt.Printf("accepted %d: %s\n", count, summarizeRequest(r))
		defer func() {
			fmt.Printf("served   %d: %s => %s\n", count, summarizeRequest(r), summarizeResponse(rw, t))
			count++
		}()
		h.ServeHTTP(rw, r)
	})
}
