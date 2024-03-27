package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"rams/pkg/models"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) error {
	user.CreatedAt = time.Now()

	query := fmt.Sprintf(`
	insert into %v (
		full_name,
		email,
		login,
		password,
		created_at
	) values (
		:full_name,
		:email,
		:login,
		:password,
		:created_at
	) returning id`, models.UsersTable)

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err = tx.QueryRow(query, user).Scan(&user.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetUser(ctx context.Context, username, password string) (user models.User, err error) {
	query := fmt.Sprintf("SELECT  id, full_name, login, email, organization, is_removed FROM %s WHERE login=$1 AND password=$2", models.UsersTable)
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return user, err
	}
	defer tx.Rollback()

	if err = tx.Get(&user, query, username, password); err != nil {
		return user, err
	}

	return user, tx.Commit()
}

func (r *Repository) GetUserByID(ctx context.Context, user models.User) (models.User, error) {
	query := fmt.Sprintf("SELECT id, full_name, login, email FROM %s WHERE id=$1", models.UsersTable)
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return user, err
	}
	defer tx.Rollback()

	if err = tx.Get(&user, query, user.ID); err != nil {
		return user, err
	}

	return user, tx.Commit()
}

func (r *Repository) GetListUser(ctx context.Context) (users []models.User, err error) {
	query := fmt.Sprintf("SELECT id, email, is_removed, full_name, login FROM %s order by id desc", models.UsersTable)
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return users, err
	}
	defer tx.Rollback()

	if err = r.db.Select(&users, query); err != nil {
		return users, err
	}

	return users, tx.Commit()
}

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	query := fmt.Sprintf("SELECT full_name, role, login, is_removed, organization, email  FROM %s WHERE login=$1 and is_removed=false", models.UsersTable)
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return user, err
	}
	defer tx.Rollback()

	if err = tx.Get(&user, query, login); err != nil {
		return user, err
	}

	return user, tx.Commit()
}
