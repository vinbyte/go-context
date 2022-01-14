package repository

import (
	"context"
	"database/sql"
)

type User struct {
	Name  string
	Email string
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) GetUser(ctx context.Context) (users []User, err error) {
	query := "SELECT name, email FROM user"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
