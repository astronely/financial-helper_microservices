package converter

import (
	"github.com/astronely/financial-helper_microservices/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.Info) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
	}
}
