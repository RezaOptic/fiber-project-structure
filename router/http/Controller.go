package httpEngine

import (
	"github.com/RezaOptic/fiber-project-structure/logic"
	"github.com/RezaOptic/fiber-project-structure/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// Interface
type UserControllerInterface interface {
	SubmitNewUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	EditUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

// 
type UserControllerStruct struct {
	Self      UserControllerInterface
	UserLogic logic.UserLogicInterface
}

// New
func NewUserController(ul logic.UserLogicInterface) UserControllerInterface {
	x := &UserControllerStruct{UserLogic: ul}
	x.Self = x
	return x
}

func (u *UserControllerStruct) SubmitNewUser(c *fiber.Ctx) error {
	var ReqData model.SubmitUserRequest
	err := c.BodyParser(&ReqData)
	if err != nil {
		return err
	}
	user, err := u.UserLogic.SubmitNewUser(c.Context(), ReqData.Username)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserControllerStruct) GetUsers(c *fiber.Ctx) error {
	sort := "desc"
	if c.Query("sort", "0") == "0" {
		sort = "desc"
	} else {
		sort = "asc"
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return err
	}
	if page < 1 {
		page = 1
	}
	perPage, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return err
	}
	if perPage < 1 {
		perPage = 1
	}
	users, hasNextPage, err := u.UserLogic.GetUsers(c.Context(), sort, page, perPage)
	if err != nil {
		return err
	}
	return c.JSON(model.ResponseUsers{
		Users:       users,
		HasNextPage: hasNextPage,
	})
}

func (u *UserControllerStruct) GetUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id", ""))
	if err != nil {
		return err
	}
	user, err := u.UserLogic.GetUser(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserControllerStruct) EditUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id", ""))
	if err != nil {
		return err
	}
	var ReqData model.UpdateUserRequest
	err = c.BodyParser(&ReqData)
	if err != nil {
		return err
	}
	user, err := u.UserLogic.UpdateUser(c.Context(), &ReqData, userID)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserControllerStruct) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id", ""))
	if err != nil {
		return err
	}
	err = u.UserLogic.DeleteUser(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusNoContent)
}
