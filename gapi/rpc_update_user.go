package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/IkehAkinyemi/mono-finance/db/sqlc"
	"github.com/IkehAkinyemi/mono-finance/pb"
	"github.com/IkehAkinyemi/mono-finance/utils"
	"github.com/IkehAkinyemi/mono-finance/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashedPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot hash password: %s", err)
		}
		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  req.Password != nil,
		}
		arg.PasswordChangedAt = sql.NullTime{
			Time: time.Now(),
			Valid: true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user : %s", err)
	}
	resp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Email != nil {
		if err := validator.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.FullName != nil {
		if err := validator.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}

	if req.Password != nil {
		if err := validator.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	return
}
