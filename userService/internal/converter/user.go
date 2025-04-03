package converter

import (
	"github.com/astronely/financial-helper_microservices/userService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/userService/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserListFromService(users []*model.User) []*desc.User {
	var userList []*desc.User
	for _, user := range users {
		userList = append(userList, ToUserFromService(user))
	}
	return userList
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Email: info.Email,
		Name:  info.Name,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Email: info.Email,
		Name:  info.Name,
	}
}

func ToUpdateUserInfoFromDesc(info *desc.UpdateRequest) *model.UpdateUserInfo {
	return &model.UpdateUserInfo{
		ID:       info.Id,
		Email:    info.GetInfo().GetEmail().GetValue(),
		Name:     info.GetInfo().GetName().GetValue(),
		Password: info.GetInfo().GetPassword().GetValue(),
	}
}
