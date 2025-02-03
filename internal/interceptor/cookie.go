package interceptor

import (
	"context"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func MetadataToCookieInterceptor(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	if md, ok := runtime.ServerMetadataFromContext(ctx); ok {
		cookies := md.HeaderMD.Get("set-cookie")
		for _, cookie := range cookies {
			logger.Debug("Cookies",
				"cookie", cookie,
			)
			w.Header().Add("Set-Cookie", cookie)
		}
		md.HeaderMD.Delete("set-cookie")
	}
	return nil
}
