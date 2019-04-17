package middlewares

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CreateLoggingMiddleware(writer io.Writer, prefix string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lw := NewLogResWriter(w)
			start := time.Now()
			handler.ServeHTTP(lw, r)

			res := []byte(fmt.Sprintf("[%s] %s %s %s %d %s %s\n", prefix, r.Method, r.RemoteAddr, r.URL.Path,
				lw.statusCode, http.StatusText(lw.statusCode), time.Since(start)))
			_, _ = writer.Write(res)
		})
	}

}

func CreateCorsMiddleware(allowedHosts []string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		c := cors.New(cors.Options{
			AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Charset", "Content-Type", "Access-Control-Allow-Credentials"},
			AllowedOrigins:     allowedHosts,
			AllowCredentials:   true,
			AllowedMethods:     []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
			OptionsPassthrough: true,
			Debug:              false,
		})
		return c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			handler.ServeHTTP(w, r)
		}))
	}
}

func CreatePanicRecoveryMiddleware() mux.MiddlewareFunc {
	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))
}

func CreateCheckAuthMiddleware(secret []byte, cookieField string,
	checkFunc func(request *http.Request, secret []byte, cookieField string) bool) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if checkFunc(r, secret, cookieField) {
				handler.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Check auth middleware fail")
		})
	}
}
