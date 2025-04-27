package converter

import (
	"github.com/astronely/financial-helper_microservices/boardService/internal/model"
	modelRepo "github.com/astronely/financial-helper_microservices/boardService/internal/repository/pg/board/model"
)

func ToBoardFromRepo(board *modelRepo.Board) *model.Board {
	return &model.Board{
		ID:        board.ID,
		Info:      ToBoardInfoFromRepo(board.Info),
		UpdatedAt: board.UpdatedAt,
		CreatedAt: board.CreatedAt,
	}
}

func ToBoardInfoFromRepo(boardInfo modelRepo.BoardInfo) model.BoardInfo {
	return model.BoardInfo{
		Name:        boardInfo.Name,
		Description: boardInfo.Description,
		OwnerID:     boardInfo.OwnerID,
	}
}

func ToBoardListFromRepo(boards []*modelRepo.Board) []*model.Board {
	var boardList []*model.Board
	for _, board := range boards {
		boardList = append(boardList, ToBoardFromRepo(board))
	}
	return boardList
}

func ToBoardUserFromRepo(boardUser *modelRepo.BoardUser) *model.BoardUser {
	return &model.BoardUser{
		BoardID:   boardUser.BoardID,
		UserID:    boardUser.UserID,
		Role:      boardUser.Role,
		CreatedAt: boardUser.CreatedAt,
	}
}

func ToBoardUserListFromRepo(boardUsers []*modelRepo.BoardUser) []*model.BoardUser {
	var boardUserList []*model.BoardUser
	for _, boardUser := range boardUsers {
		boardUserList = append(boardUserList, ToBoardUserFromRepo(boardUser))
	}
	return boardUserList
}
