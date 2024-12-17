package controllers

import (
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/pilinux/argon2"
)

func LoginUsers(ctx *gin.Context) {
	var formUser Users
	ctx.ShouldBind(&formUser)

	foundUser := models.FindOneUserByEmail(formUser.Email)
	if foundUser == (models.Users{}) {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Not Found",
		})
		return
	}

	match, _ := argon2.ComparePasswordAndHash(formUser.Password, formUser.Password, foundUser.Password)
	if !match {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Password",
		})
		return
	}

	token := lib.GeneratedToken(struct {
		UserId int `json:"userId"`
	}{
		UserId: foundUser.Id,
	})
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Login Success",
		Results: token,
	})

}

func RegisterUser(ctx *gin.Context) {
	var formUser Users
	ctx.ShouldBind(&formUser)

	for _, getuser := range Listusers {
		if formUser.Email == getuser.Email {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email Has Registered",
			})
		}
		if !strings.Contains(formUser.Email, "@") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include @",
			})
			return
		}
		if len(formUser.Password) < 6 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email Must be 6 Character",
			})
			return
		}
		if !strings.ContainsAny(formUser.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include Uppercase Character",
			})
			return
		}
		if !strings.ContainsAny(formUser.Password, "0123456789") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include One Number",
			})
			return
		}
		if !strings.ContainsAny(formUser.Password, "!@#$%^&*()-_=+[]{}|;:,.<>?") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include Unique Character",
			})
			return
		}
	}
	hash, _ := argon2.CreateHash(formUser.Password, formUser.Password, argon2.DefaultParams)

	var row = Users{
		Email:    formUser.Email,
		Password: hash,
	}

	if formUser.Email == "" {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Fill Your Email",
		})
	} else {
		row.Id = len(Listusers) + 1
		Listusers = append(Listusers, row)
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register Success",
			Results: row,
		})

		return
	}

}
