package converter

import (
	"github.com/astronely/financial-helper_microservices/internal/model"
	"github.com/astronely/financial-helper_microservices/internal/repository/user/modelRepo"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	//var updatedAt *timestamppb.Timestamp
	//if user.UpdatedAt.Valid {
	//	updatedAt = timestamppb.New(user.UpdatedAt.Time)
	//}

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
