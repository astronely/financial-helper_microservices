package board

import (
	"github.com/astronely/financial-helper_microservices/boardService/internal/service"
	desc "github.com/astronely/financial-helper_microservices/boardService/pkg/board_v1"
)

type Implementation struct {
	desc.UnimplementedBoardV1Server
	service service.BoardService
}

func NewImplementation(service service.BoardService) *Implementation {
	return &Implementation{
		service: service,
	}
}
