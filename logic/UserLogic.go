package logic

import (
	"context"
	"github.com/RezaOptic/fiber-project-structure/model"
	"github.com/RezaOptic/fiber-project-structure/repository"
)

// Interface
type UserLogicInterface interface {
	SubmitNewUser(c context.Context, username string) (*model.User, error)
	GetUsers(c context.Context, sortType string, page, perPage int) ([]model.User, bool, error)
	GetUser(c context.Context, userID int) (*model.User, error)
	UpdateUser(c context.Context, newUserData *model.UpdateUserRequest, userID int) (*model.User, error)
	DeleteUser(c context.Context, userID int) error
}

// 
type UserLogic struct {
	Self     UserLogicInterface
	UserRepo repository.UserRepoInterface
}

// New
func NewUserLogic(ur *repository.UserRepoInterface) UserLogicInterface {
	x := &UserLogic{UserRepo: *ur}
	x.Self = x
	return x
}

func (u *UserLogic) SubmitNewUser(c context.Context, username string) (*model.User, error) {
	user, err := u.UserRepo.SubmitNewUser(c, username)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserLogic) GetUsers(c context.Context, sortType string, page, perPage int) ([]model.User, bool, error) {
	hasNextPage := false
	page--
	perPage++
	users, err := u.UserRepo.GetUsers(c, sortType, page, perPage)
	if err != nil {
		return nil, hasNextPage, err
	}
	if len(users) == perPage {
		hasNextPage = true
		users = users[:len(users)-1]
	}
	return users, hasNextPage, err
}

func (u *UserLogic) GetUser(c context.Context, userID int) (*model.User, error) {
	user, err := u.UserRepo.GetUser(c, userID)
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *UserLogic) UpdateUser(c context.Context, newUserData *model.UpdateUserRequest, userID int) (*model.User, error) {
	user, err := u.UserRepo.EditUser(c, newUserData, userID)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserLogic) DeleteUser(c context.Context, userID int) error {
	err := u.UserRepo.DeleteUser(c, userID)
	if err != nil {
		return err
	}
	return err
}
