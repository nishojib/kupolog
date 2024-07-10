package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/uptrace/bun"
)

func (r *Repository) InsertAndLinkAccount(
	ctx context.Context,
	user *user.User,
	account *user.Account,
) error {
	err := r.db.RunInTx(
		ctx,
		&sql.TxOptions{},
		func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewInsert().Model(user).Exec(ctx)
			if err != nil {
				// If the error is a unique constraint error, try to fetch the user by email
				if strings.Contains(
					err.Error(),
					"SQLite error: UNIQUE constraint failed: users.email",
				) {
					err = tx.NewSelect().Model(user).Where("email = ?", user.Email).Scan(ctx)
					if err != nil {
						return err
					}
				} else {
					return err
				}
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

func (r *Repository) GetUserByProviderID(
	ctx context.Context,
	providerAccountID string,
) (user.User, error) {
	var u user.User
	err := r.db.NewSelect().
		Model(&u).
		Join("JOIN accounts ON accounts.user_id = user.id").
		Where("accounts.provider_account_id = ?", providerAccountID).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return user.User{}, repoErrors.ErrRecordNotFound
		default:
			return user.User{}, err
		}
	}

	return u, nil
}

func (r *Repository) UpdateUser(ctx context.Context, u *user.User) error {
	result, err := r.db.NewUpdate().
		Model(u).
		Set("title = ?", u.Name).
		Set("author = ?", u.Email).
		Set("version = ?", u.Version+1).
		Where("id = ?", u.ID).
		Where("version = ?", u.Version).
		Exec(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return repoErrors.ErrEditConflict
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repoErrors.ErrEditConflict
	}

	return nil
}