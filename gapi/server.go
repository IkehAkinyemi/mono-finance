package gapi

import (
	"fmt"

	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	"github.com/IkehAkinyemi/mono-finance/pb"
	"github.com/IkehAkinyemi/mono-finance/token"
	"github.com/IkehAkinyemi/mono-finance/utils"
	"github.com/IkehAkinyemi/mono-finance/worker"
)

// A Server serves gRPC requests for the banking system
type Server struct {
	pb.UnimplementedMonoFinanceServer
	config          utils.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer create a new gRPC server and setup routing.
func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
