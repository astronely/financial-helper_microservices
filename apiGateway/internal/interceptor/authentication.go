package interceptor

import (
	"context"
	accessService "github.com/astronely/financial-helper_microservices/apiGateway/pkg/access_v1"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
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

			//authHeader := r.Header.Get("Authorization")

			ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
			defer cancel()

			//if authHeader != "" {
			//	md := metadata.Pairs("authorization", authHeader)
			//	ctx = metadata.NewOutgoingContext(ctx, md)
			//}

			req := &accessService.CheckRequest{
				EndpointAddress: r.URL.Path,
			}

			var header, trailer metadata.MD

			res, err := accessClient.Check(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				http.Error(w, "AuthService error: "+err.Error(), http.StatusUnauthorized)
				//rw.WriteHeader(http.StatusUnauthorized)
				rw.headers.Set("X-Error", err.Error())
				logger.Error("AuthService error",
					"error", err.Error(),
				)
				return
			}

			for _, v := range header.Get("set-cookie") {
				rw.headers.Add("Set-Cookie", v)
			}

			existingMD, _ := metadata.FromOutgoingContext(r.Context())
			//logger.Debug("existing md",
			//	"md", existingMD)

			if len(header.Get("set-cookie")) > 0 {
				var newContext context.Context
				newAccessToken := header.Get("set-cookie")[0]
				newAccessToken = strings.Split(newAccessToken, ";")[0]
				newAccessToken = strings.Split(newAccessToken, "=")[1]

				newRefreshToken := header.Get("set-cookie")[1]
				newRefreshToken = strings.Split(newRefreshToken, ";")[0]
				newRefreshToken = strings.Split(newRefreshToken, "=")[1]

				//logger.Debug("New Tokens",
				//	"newAccessToken", newAccessToken,
				//	"newRefreshToken", newRefreshToken,
				//)

				newMD := existingMD.Copy()
				newMD.Set(accessTokenName, newAccessToken)
				newMD.Set(refreshTokenName, newRefreshToken)
				newContext = metadata.NewOutgoingContext(r.Context(), newMD)

				r = r.WithContext(newContext)

				//newC, _ := metadata.FromOutgoingContext(r.Context())
				//logger.Debug("New md",
				//	"md", newC)
			}

			//logger.Debug("AuthService",
			//	"response", res,
			//	"metadata", header,
			//	"trailer", trailer,
			//)

			if !res.IsAllowed {
				http.Error(rw, "Unauthorized", http.StatusUnauthorized)
				//rw.WriteHeader(http.StatusUnauthorized)
				logger.Error("Unauthorized check",
					"error", err.Error(),
				)
				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}
