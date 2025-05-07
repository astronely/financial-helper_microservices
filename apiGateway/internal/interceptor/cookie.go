package interceptor

import (
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"net/http"
	"strings"
)

func CookieInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw, ok := w.(*responseWrapper)
		if !ok {
			panic("invalid ResponseWriter type: expected responseWrapper")
		}
		//logger.Debug("first in cookie")
		next.ServeHTTP(w, r)
		//logger.Debug("second in cookie")

		// Заменяем grpc-set-cookie на set-cookie
		for k, vv := range rw.headers {
			if strings.ToLower(k) == "grpc-metadata-set-cookie" {
				for _, v := range vv {
					rw.headers.Add("Set-Cookie", v)
				}
				delete(rw.headers, k)
				rw.ResponseWriter.Header().Del(k)
			}
		}

		for k, vv := range rw.headers {
			for _, v := range vv {
				rw.ResponseWriter.Header().Add(k, v)
			}
		}

		logger.Debug("Cookies",
			"header", w.Header())
	})
}

type responseWrapper struct {
	http.ResponseWriter
	headers http.Header
	status  int
	written bool
}

func (h *responseWrapper) Header() http.Header {
	return h.headers
}

func (h *responseWrapper) WriteHeader(statusCode int) {
	if h.written {
		return
	}
	h.status = statusCode
	h.written = true
}

func (h *responseWrapper) Write(b []byte) (int, error) {
	if !h.written {
		h.WriteHeader(http.StatusOK)
	}
	h.copyHeaders()
	h.ResponseWriter.WriteHeader(h.status)
	return h.ResponseWriter.Write(b)
}

func (h *responseWrapper) copyHeaders() {
	for k, vv := range h.headers {
		for _, v := range vv {
			h.ResponseWriter.Header().Add(k, v)
		}
	}
}
