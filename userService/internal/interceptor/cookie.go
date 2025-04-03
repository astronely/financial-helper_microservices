package interceptor

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func MetadataToCookieInterceptor(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	if md, ok := runtime.ServerMetadataFromContext(ctx); ok {
		cookies := md.HeaderMD.Get("set-cookie")
		for _, cookie := range cookies {
			//logger.Debug("Cookies",
			//	"cookie", cookie,
			//)
			w.Header().Add("Set-Cookie", cookie)
		}
		authHeader := md.HeaderMD.Get("authorization")
		if len(authHeader) > 0 {
			//logger.Debug("Authorization",
			//	"header", authHeader,
			//)
			w.Header().Add("Authorization", authHeader[0])
			//logger.Debug("Auth header:",
			//	"value: ", w.Header().Values("Authorization"))
		}
	}
	return nil
}
