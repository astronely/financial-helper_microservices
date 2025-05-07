package interceptor

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	accessTokenName  = "token"
	refreshTokenName = "refreshToken"
)

func SetCookiesInterceptor(ctx context.Context) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw, ok := w.(*responseWrapper)
			if !ok {
				panic("invalid ResponseWriter type: expected responseWrapper")
			}

			cookies := r.Cookies()

			var token, refreshToken string
			for _, cookie := range cookies {
				if cookie.Name == accessTokenName {
					token = cookie.Value
				}
				if cookie.Name == refreshTokenName {
					refreshToken = cookie.Value
				}
			}

			//logger.Debug("Cookies",
			//	"cookies", cookies,
			//	"token", token,
			//)

			ctxWithMetadata := metadata.AppendToOutgoingContext(r.Context(), accessTokenName, token)
			ctxWithMetadata = metadata.AppendToOutgoingContext(ctxWithMetadata, refreshTokenName, refreshToken)
			r = r.WithContext(ctxWithMetadata)
			//logger.Debug("first in logger")

			next.ServeHTTP(rw, r)

			//logger.Debug("second in logger")
		})
	}
}

//func SetCookiesInterceptor(ctx context.Context) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			//logger.Debug("Inside AuthInterceptor",
//			//	"method", r.Method,
//			//	"url", r.URL.String(),
//			//)
//			rw, ok := w.(*responseWrapper)
//			if !ok {
//				panic("invalid ResponseWriter type: expected responseWrapper")
//			}
//			cookies := r.Cookies()
//			logger.Debug("Cookies",
//				"cookies", cookies,
//				)
//			//authHeader := r.Header.Get("Authorization")
//			//
//			//ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
//			//defer cancel()
//			//
//			//if authHeader != "" {
//			//	md := metadata.Pairs("authorization", authHeader)
//			//	ctx = metadata.NewOutgoingContext(ctx, md)
//			//}
//			//
//			//req := &accessService.CheckRequest{
//			//	EndpointAddress: r.URL.Path,
//			//}
//			//
//			//res, err := accessClient.Check(ctx, req)
//			//if err != nil {
//			//	http.Error(w, "AuthService error: "+err.Error(), http.StatusUnauthorized)
//			//	rw.WriteHeader(http.StatusUnauthorized)
//			//	rw.headers.Set("X-Error", err.Error())
//			//	logger.Error("AuthService error",
//			//		"error", err.Error(),
//			//	)
//			//	return
//			//}
//			//
//			//if !res.IsAllowed {
//			//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			//	rw.WriteHeader(http.StatusUnauthorized)
//			//	logger.Error("Unauthorized check",
//			//		"error", err.Error(),
//			//	)
//			//	return
//			//}
//
//			next.ServeHTTP(rw, r)
//		})
//	}
//}
