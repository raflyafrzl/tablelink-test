package repository

import (
	"context"
	user2 "test_tablelink/user"
)

func (r *Repository) FindUserByUsername(ctx context.Context, username string) (*user2.User, error) {
	query := `SELECT id, xid, username, password, email FROM users WHERE username = ?`
	query = r.Db.Rebind(query)

	var user user2.User
	err := r.Db.GetContext(ctx, &user, query, username)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *user2.User) error {

	query := `INSERT INTO users (xid, username, password, email) VALUES (?, ?, ?, ?)`

	query = r.Db.Rebind(query)
	_, err := r.Db.ExecContext(ctx, query, user.Xid, user.Username, user.Password, user.Email)

	if err != nil {
		return err
	}

	return nil
}
