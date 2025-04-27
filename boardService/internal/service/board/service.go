package board

import (
	"github.com/astronely/financial-helper_microservices/boardService/internal/repository"
	def "github.com/astronely/financial-helper_microservices/boardService/internal/service"
)

var _ def.BoardService = (*serv)(nil)

type serv struct {
	boardRepository      repository.BoardRepository
	boardRedisRepository repository.BoardRedisRepository
}

func NewService(boardRepository repository.BoardRepository, boardRedisRepository repository.BoardRedisRepository) def.BoardService {
	return &serv{
		boardRepository:      boardRepository,
		boardRedisRepository: boardRedisRepository,
	}
}
