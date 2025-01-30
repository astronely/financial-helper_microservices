package auth

import (
	"context"
	"github.com/astronely/financial-helper_microservices/pkg/logger"
)

func (s *serv) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	logger.Debug("User in service layer, login",
		"ID", user.ID,
	)
	//if !utils.VerifyPassword(user.Info.Password, password) {
	//	return "", errors.New("invalid password")
	//}

	return "test success", nil
}
