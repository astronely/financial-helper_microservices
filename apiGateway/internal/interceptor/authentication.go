package interceptor

import (
	"context"
	accessService "github.com/astronely/financial-helper_microservices/apiGateway/pkg/access_v1"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

func AuthInterceptor(ctx context.Context, accessClient accessService.AccessV1Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//logger.Debug("Inside AuthInterceptor",
			//	"method", r.Method,
			//	"url", r.URL.String(),
			//)
			rw, ok := w.(*responseWrapper)
			if !ok {
				panic("invalid ResponseWriter type: expected responseWrapper")
			}

			authHeader := r.Header.Get("Authorization")

			ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
			defer cancel()

			if authHeader != "" {
				md := metadata.Pairs("authorization", authHeader)
				ctx = metadata.NewOutgoingContext(ctx, md)
			}

			req := &accessService.CheckRequest{
				EndpointAddress: r.URL.Path,
			}

			res, err := accessClient.Check(ctx, req)
			if err != nil {
				http.Error(w, "AuthService error: "+err.Error(), http.StatusUnauthorized)
				rw.WriteHeader(http.StatusUnauthorized)
				rw.headers.Set("X-Error", err.Error())
				logger.Error("AuthService error",
					"error", err.Error(),
				)
				return
			}

			if !res.IsAllowed {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				rw.WriteHeader(http.StatusUnauthorized)
				logger.Error("Unauthorized check",
					"error", err.Error(),
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
