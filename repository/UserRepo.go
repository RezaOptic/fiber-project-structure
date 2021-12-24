package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RezaOptic/fiber-project-structure/model"
	"github.com/RezaOptic/fiber-project-structure/repository/sqlQuery"
	"time"
)

type UserRepoInterface interface {
	SubmitNewUser(ctx context.Context, username string) (*model.User, error)
	GetUsers(ctx context.Context, sortType string, page, perPage int) ([]model.User, error)
	GetUser(ctx context.Context, userID int) (*model.User, error)
	EditUser(ctx context.Context, newUserData *model.UpdateUserRequest, userID int) (*model.User, error)
	DeleteUser(ctx context.Context, userID int) error
}

type UserRepo struct {
	Self         UserRepoInterface
	DBConnection *sql.Conn
}

func NewUserRepo(s *sql.Conn) UserRepoInterface {
	x := &UserRepo{DBConnection: s}
	x.Self = x
	return x
}

func (u *UserRepo) SubmitNewUser(ctx context.Context, username string) (*model.User, error) {
	User := model.User{
		Username:  username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	lastInsertId := 0
	err := u.DBConnection.QueryRowContext(ctx, sqlQuery.InsertNewUserQuery, User.Username,
		User.CreatedAt, User.UpdatedAt).Scan(&lastInsertId)
	if err != nil {
		return &User, err
	}
	User.ID = lastInsertId
	return &User, nil
}

func (u *UserRepo) GetUsers(ctx context.Context, sortType string, page, perPage int) ([]model.User, error) {
	var users []model.User
	query := fmt.Sprintf(sqlQuery.GetUsersQuery, sortType, perPage, perPage-1*page)
	rows, err := u.DBConnection.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u *UserRepo) GetUser(ctx context.Context, userID int) (*model.User, error) {
	var user model.User
	err := u.DBConnection.QueryRowContext(ctx, sqlQuery.GetUserQuery, userID).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &user, err
	}
	return &user, err
}

func (u *UserRepo) EditUser(ctx context.Context, newUserData *model.UpdateUserRequest, userID int) (*model.User, error) {
	User := model.User{
		ID:        userID,
		Username:  newUserData.Username,
		UpdatedAt: time.Now().UTC(),
	}
	_, err := u.DBConnection.ExecContext(ctx, sqlQuery.UpdateUserQuery, User.Username, User.UpdatedAt, userID)
	if err != nil {
		return &User, err
	}
	return &User, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, userID int) error {
	_, err := u.DBConnection.ExecContext(ctx, sqlQuery.DeleteUserQuery, userID)
	if err != nil {
		return err
	}
	return nil
}
