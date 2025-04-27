package model

type InvitationData struct {
	BoardID int64  `json:"board_id"`
	UserID  int64  `json:"user_id"`
	Role    string `json:"role"`
}
