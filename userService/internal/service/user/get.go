package user

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/pkg/logger"
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	"github.com/astronely/financial-helper_microservices/userService/internal/utils"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var userID int64
	var err error

	if id > 0 {
		userID = id
	} else {
		userID, err = utils.GetUserIdFromContext(ctx, s.tokenConfig.AccessTokenKey())
		if err != nil {
			logger.Error("cannot get userID from context",
				"error", err.Error(),
			)
			return nil, err
		}
	}

	user, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
