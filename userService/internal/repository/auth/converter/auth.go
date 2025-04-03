package converter

import (
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/userService/internal/repository/auth/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.UserAuth {
	return &model.UserAuth{
		ID:   user.ID,
		Info: ToUserInfoFromRepo(user.Info),
	}
}

func ToUserInfoFromRepo(info modelRepo.Info) model.UserAuthInfo {
	return model.UserAuthInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
	}
}
