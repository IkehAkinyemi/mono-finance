package db

import (
	"context"
	"database/sql"
)

// A VerifyEmailTxParams contains the input parameters of the verify email transaction.
type VerifyEmailTxParams struct {
	EmailId int64
	SecretCode string
}

// A VerifyEmailTxResult contains the result of the verify email transaction.
type VerifyEmailTxResult struct {
	User User
	VerifyEmail VerifyEmail
}

// VerifyEmailTx performs a verify email transaction
func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID: arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: sql.NullBool{
				Bool: true,
				Valid: true,
			},
		})

		return err
	})

	return result, err
}
