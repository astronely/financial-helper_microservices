package env

import (
	"errors"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/config"
	"net"
	"os"
)

const (
	grpcUserHostEnvName = "GRPC_USER_SERVICE_HOST"
	grpcUserPortEnvName = "GRPC_USER_SERVICE_PORT"

	grpcAuthHostEnvName = "GRPC_AUTH_SERVICE_HOST"
	grpcAuthPortEnvName = "GRPC_AUTH_SERVICE_PORT"

	grpcFinanceHostEnvName = "GRPC_FINANCE_SERVICE_HOST"
	grpcFinancePortEnvName = "GRPC_FINANCE_SERVICE_PORT"

	grpcNoteHostEnvName = "GRPC_NOTE_SERVICE_HOST"
	grpcNotePortEnvName = "GRPC_NOTE_SERVICE_PORT"

	grpcBoardHostEnvName = "GRPC_BOARD_SERVICE_HOST"
	grpcBoardPortEnvName = "GRPC_BOARD_SERVICE_PORT"
)

type grpcConfig struct {
	userHost string
	userPort string

	authHost string
	authPort string

	financeHost string
	financePort string

	noteHost string
	notePort string

	boardHost string
	boardPort string
}

func NewGRPCConfig() (config.GRPCConfig, error) {
	userHost := os.Getenv(grpcUserHostEnvName)
	if len(userHost) == 0 {
		return nil, errors.New("userHost not found in env")
	}
	userPort := os.Getenv(grpcUserPortEnvName)
	if len(userPort) == 0 {
		return nil, errors.New("userPort not found in env")
	}

	authHost := os.Getenv(grpcAuthHostEnvName)
	if len(authHost) == 0 {
		return nil, errors.New("authHost not found in env")
	}
	authPort := os.Getenv(grpcAuthPortEnvName)
	if len(authPort) == 0 {
		return nil, errors.New("authPort not found in env")
	}

	financeHost := os.Getenv(grpcFinanceHostEnvName)
	if len(financeHost) == 0 {
		return nil, errors.New("financeHost not found in env")
	}
	financePort := os.Getenv(grpcFinancePortEnvName)
	if len(financePort) == 0 {
		return nil, errors.New("financePort not found in env")
	}

	noteHost := os.Getenv(grpcNoteHostEnvName)
	if len(noteHost) == 0 {
		return nil, errors.New("noteHost not found in env")
	}
	notePort := os.Getenv(grpcNotePortEnvName)
	if len(notePort) == 0 {
		return nil, errors.New("notePort not found in env")
	}

	boardHost := os.Getenv(grpcBoardHostEnvName)
	if len(boardHost) == 0 {
		return nil, errors.New("boardHost not found in env")
	}
	boardPort := os.Getenv(grpcBoardPortEnvName)
	if len(boardPort) == 0 {
		return nil, errors.New("boardPort not found in env")
	}

	return &grpcConfig{
		userHost: userHost,
		userPort: userPort,

		authHost: authHost,
		authPort: authPort,

		financeHost: financeHost,
		financePort: financePort,

		noteHost: noteHost,
		notePort: notePort,

		boardHost: boardHost,
		boardPort: boardPort,
	}, nil
}

func (c *grpcConfig) UserAddress() string {
	return net.JoinHostPort(c.userHost, c.userPort)
}

func (c *grpcConfig) AuthAddress() string {
	return net.JoinHostPort(c.authHost, c.authPort)
}

func (c *grpcConfig) FinanceAddress() string {
	return net.JoinHostPort(c.financeHost, c.financePort)
}

func (c *grpcConfig) NoteAddress() string {
	return net.JoinHostPort(c.noteHost, c.notePort)
}

func (c *grpcConfig) BoardAddress() string {
	return net.JoinHostPort(c.boardHost, c.boardPort)
}
