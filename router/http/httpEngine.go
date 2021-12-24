package httpEngine

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Run(Port string, UserController UserControllerInterface) {
	engine := fiber.New()
	V1 := engine.Group("v1")
	V1.Post("/users", UserController.SubmitNewUser)
	V1.Get("/users", UserController.GetUsers)
	V1.Get("/users/:user_id", UserController.GetUser)
	V1.Put("/users/:user_id", UserController.EditUser)
	V1.Delete("/users/:user_id", UserController.DeleteUser)
	fmt.Println(engine.Listen(fmt.Sprintf(":%s", Port)))
}
