package gapi

import (
	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	"github.com/IkehAkinyemi/mono-finance/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		Email: user.Email,
		FullName: user.FullName,
		PasswordChangeAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}