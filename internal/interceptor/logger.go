package interceptor

import (
	"context"
	"fmt"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
	"google.golang.org/grpc"
	"time"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(),
			"method", info.FullMethod,
			"req", req,
		)
		return nil, err
	}

	logger.Info("request",
		"method", info.FullMethod,
		"req", req,
		"res", res,
		"duration", fmt.Sprintf("%f seconds", time.Since(now).Seconds()),
	)

	return res, err
}
