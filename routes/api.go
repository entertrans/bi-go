package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Api() {
	auth := controllers.AuthController{}

	facades.Route().Post("/register", auth.Register)
	facades.Route().Post("/login", auth.Login)
}
