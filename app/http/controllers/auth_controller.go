package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/hash"
	"github.com/goravel/framework/support/jwt"
)

type AuthController struct{}

func (r *AuthController) Register(ctx http.Context) {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	hashedPassword, _ := hash.Make(password)

	user := map[string]interface{}{
		"name":     ctx.Request().Input("name"),
		"email":    email,
		"password": hashedPassword,
	}

	result := facades.Orm().Query().Table("users").Create(user)
	if result.Error != nil {
		ctx.JSON(500, map[string]interface{}{
			"message": "Gagal register",
		})
		return
	}

	ctx.JSON(201, map[string]interface{}{
		"message": "User berhasil terdaftar",
	})
}

func (r *AuthController) Login(ctx http.Context) {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	var user map[string]interface{}
	result := facades.Orm().Query().Table("users").Where("email", email).First(&user)
	if result.Error != nil {
		ctx.JSON(401, map[string]interface{}{
			"message": "User tidak ditemukan",
		})
		return
	}

	if !hash.Check(password, user["password"].(string)) {
		ctx.JSON(401, map[string]interface{}{
			"message": "Password salah",
		})
		return
	}

	token, _ := jwt.Make(map[string]interface{}{
		"id":    user["id"],
		"email": user["email"],
	})

	ctx.JSON(200, map[string]interface{}{
		"message": "Login berhasil",
		"token":   token,
	})
}
