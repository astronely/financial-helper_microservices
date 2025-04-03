package interceptor

import (
	_ "github.com/astronely/financial-helper_microservices/userService/pkg/logger"
)

//type AuthInterceptor struct {
//	accessService service.AccessService
//}
//
//func NewAuthInterceptor(accessService service.AccessService) *AuthInterceptor {
//	return &AuthInterceptor{
//		accessService: accessService,
//	}
//}
//
//func (a *AuthInterceptor) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//
//	isAllowed, err := a.accessService.Check(ctx, info.FullMethod)
//	if err != nil {
//		return nil, err
//	}
//
//	if !isAllowed {
//		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
//	}
//
//	return handler(ctx, req)
//}
