package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/IkehAkinyemi/mono-finance/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestUpdateUserFullNameOnly(t *testing.T) {
		oldUser := createRandomUser(t)

		newFullName := utils.RandomOwner()

		updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
			Username: oldUser.Username,
			FullName: sql.NullString{
				String: newFullName,
				Valid: true,
			},
		})
		require.NoError(t, err)
		require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
		require.Equal(t, newFullName, updatedUser.FullName)
		require.Equal(t, oldUser.Email, updatedUser.Email)
		require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserEmailOnly(t *testing.T) {
		oldUser := createRandomUser(t)

		newEmail := utils.RandomEmail()

		updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
			Username: oldUser.Username,
			Email: sql.NullString{
				String: newEmail,
				Valid: true,
			},
		})
		require.NoError(t, err)
		require.NotEqual(t, oldUser.Email, updatedUser.Email)
		require.Equal(t, newEmail, updatedUser.Email)
		require.Equal(t, oldUser.FullName, updatedUser.FullName)
		require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserPasswordOnly(t *testing.T) {
		oldUser := createRandomUser(t)

		newHashedPassword, err := utils.HashedPassword(utils.RandomString(10))
		require.NoError(t, err)

		updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
			Username: oldUser.Username,
			HashedPassword: sql.NullString{
				String: newHashedPassword,
				Valid: true,
			},
		})
		require.NoError(t, err)
		require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
		require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
		require.Equal(t, oldUser.FullName, updatedUser.FullName)
		require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
		oldUser := createRandomUser(t)

		newFullName := utils.RandomOwner()
		newEmail := utils.RandomEmail()
		newHashedPassword, err := utils.HashedPassword(utils.RandomString(10))
		require.NoError(t, err)

		updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
			Username: oldUser.Username,
			HashedPassword: sql.NullString{
				String: newHashedPassword,
				Valid: true,
			},
			FullName: sql.NullString{
				String: newFullName,
				Valid: true,
			},
			Email: sql.NullString{
				String: newEmail,
				Valid: true,
			},
		})
		require.NoError(t, err)
		require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
		require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
		require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
		require.Equal(t, newFullName, updatedUser.FullName)
		require.NotEqual(t, oldUser.Email, updatedUser.Email)
		require.Equal(t, newEmail, updatedUser.Email)
}
