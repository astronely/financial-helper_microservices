package converter

import (
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

func ToGenerateInviteFromDesc(req *desc.GenerateInviteRequest) *model.GenerateInviteInfo {
	return &model.GenerateInviteInfo{
		BoarID: req.GetBoardId(),
		UserID: req.GetUserId(),
		Role:   req.GetRole(),
	}
}

func ToJoinInfoFromDesc(req *desc.JoinRequest) *model.JoinInfo {
	return &model.JoinInfo{
		Token: req.GetToken(),
	}
}

func ToGenerateInvite(info *model.GenerateInviteInfo) model.GenerateInviteInfo {
	return model.GenerateInviteInfo{
		BoarID: info.BoarID,
		UserID: info.UserID,
		Role:   info.Role,
	}
}
