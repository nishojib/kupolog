package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) InsertAndLinkAccount(user *models.User, account *models.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := ur.db.RunInTx(
		ctx,
		&sql.TxOptions{},
		func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewInsert().Model(user).Exec(ctx)
			if err != nil {
				return err
			}

			account.UserID = user.ID

			fmt.Printf("%+v\n", account)

			_, err = tx.NewInsert().Model(account).Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetByProviderID(providerAccountID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	err := ur.db.NewSelect().
		Model(&user).
		Join("JOIN accounts ON accounts.user_id = user.id").
		Where("accounts.provider_account_id = ?", providerAccountID).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.User{}, ErrRecordNotFound
		default:
			return models.User{}, err
		}
	}

	return user, nil
}

func (ur *UserRepository) Update(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := ur.db.NewUpdate().
		Model(user).
		Set("title = ?", user.Name).
		Set("author = ?", user.Email).
		Set("version = ?", user.Version+1).
		Where("id = ?", user.ID).
		Where("version = ?", user.Version).
		Exec(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrEditConflict
	}

	return nil
}
