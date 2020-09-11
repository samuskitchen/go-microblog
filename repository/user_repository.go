package data

import (
	"context"
	"time"

	conn "microblog/database"
	"microblog/domain/user"
)

// UserRepository manages the operations with the database that correspond to the user model.
type UserRepository struct {
	Data *conn.Data
}

// GetAll returns all users.
func (ur *UserRepository) GetAllUser(ctx context.Context) ([]user.User, error) {
	rows, err := ur.Data.DB.QueryContext(ctx, selectAllUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var user user.User
		_ = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Picture, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	return users, nil
}

// GetOne returns one user by id.
func (ur *UserRepository) GetOne(ctx context.Context, id uint) (user.User, error) {
	row := ur.Data.DB.QueryRowContext(ctx, selectUserById, id)

	var userScan user.User
	err := row.Scan(&userScan.ID, &userScan.FirstName, &userScan.LastName, &userScan.Username, &userScan.Email, &userScan.Picture, &userScan.CreatedAt, &userScan.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return userScan, nil
}

// GetByUsername returns one user by username.
func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (user.User, error) {
	row := ur.Data.DB.QueryRowContext(ctx, selectUSerByUsername, username)

	var userScan user.User
	err := row.Scan(&userScan.ID, &userScan.FirstName, &userScan.LastName, &userScan.Username,
		&userScan.Email, &userScan.Picture, &userScan.PasswordHash, &userScan.CreatedAt, &userScan.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return userScan, nil
}

// Create adds a new user.
func (ur *UserRepository) Create(ctx context.Context, user *user.User) error {
	now := time.Now().Truncate(time.Second).Truncate(time.Millisecond).Truncate(time.Microsecond)

	if user.Picture == "" {
		user.Picture = "https://placekitten.com/g/300/300"
	}

	stmt, err := ur.Data.DB.PrepareContext(ctx, insertUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, user.FirstName, user.LastName, user.Username, user.Email,
		user.Picture, user.PasswordHash, now, now,
	)

	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a user by id.
func (ur *UserRepository) Update(ctx context.Context, id uint, u user.User) error {
	now := time.Now().Truncate(time.Second).Truncate(time.Millisecond).Truncate(time.Microsecond)

	stmt, err := ur.Data.DB.PrepareContext(ctx, updateUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.FirstName, u.LastName, u.Email, u.Picture, now, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a user by id.
func (ur *UserRepository) Delete(ctx context.Context, id uint) error {

	stmt, err := ur.Data.DB.PrepareContext(ctx, deleteUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
