package interceptor

import (
	"context"
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw, ok := w.(*responseWrapper)
			if !ok {
				panic("invalid ResponseWriter type: expected responseWrapper")
			}

			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			done := make(chan struct{})
			go func() {
				next.ServeHTTP(rw, r.WithContext(ctx))
				close(done)
			}()

			select {
			case <-done:
			case <-ctx.Done():
				http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
			}
		})
	}
}
