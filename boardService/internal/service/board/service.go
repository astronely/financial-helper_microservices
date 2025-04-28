package board

import (
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/boardService/internal/service"
	"github.com/astronely/financial-helper_microservices/userService/pkg/client/db"
)

var _ def.BoardService = (*serv)(nil)

type serv struct {
	boardRepository      repository.BoardRepository
	boardRedisRepository repository.BoardRedisRepository

	txManager db.TxManager
}

func NewService(boardRepository repository.BoardRepository, boardRedisRepository repository.BoardRedisRepository, txManager db.TxManager) def.BoardService {
	return &serv{
		boardRepository:      boardRepository,
		boardRedisRepository: boardRedisRepository,
		txManager:            txManager,
	}
}
