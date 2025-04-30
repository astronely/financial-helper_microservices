package interceptor

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"net/http"
	"time"
)

func Logger(ctx context.Context) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWrapper{
				ResponseWriter: w,
				headers:        make(http.Header),
			}
			//logger.Debug("first in logger")

			next.ServeHTTP(rw, r)

			//logger.Debug("second in logger")

			logger.Info("request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.status,
				"headers", w.Header(),
				"duration", time.Since(start),
			)
		})
	}
}
