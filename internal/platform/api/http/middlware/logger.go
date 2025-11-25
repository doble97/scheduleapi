package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware envuelve un http.Handler y registra la petición.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Creamos un ResponseWriter personalizado para capturar el código de estado.
		lw := NewLoggingResponseWriter(w)

		// Pasa el control al siguiente handler (o a tu UserHandler)
		next.ServeHTTP(lw, r)

		// Log después de que el handler ha terminado
		log.Printf(
			"[%s] %s %s %d %s",
			time.Now().Format("2006/01/02 15:04:05"),
			r.Method,
			r.RequestURI,
			lw.StatusCode, // Usamos el código de estado capturado
			time.Since(start),
		)
	})
}

// LoggingResponseWriter es un wrapper para capturar el código de estado
type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// Por defecto, asumimos 200 OK si no se establece explícitamente.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
