package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"test/lib"

	"github.com/gin-gonic/gin"
	"github.com/pilinux/argon2"
)

type Users struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var Listusers = []Users{
	{
		Id:       1,
		Fullname: "Jonathan Born",
		Email:    "jonathan@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$7YMAzBLKe2AQP2kGWjrtJg$qe8VKWQc9vaGEeskvGA7A/qyTRhwEjo6gHt7v+TkA1g",
	},
	{
		Id:       2,
		Fullname: "Hanny Mony",
		Email:    "hanny@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$312xYQgejPmY6xru1NBEEA$ttmXf2fuK6Vdej4QeazjtTT8553PRTqk0lJ8dBTaUlE",
	},
	{
		Id:       3,
		Fullname: "Saitoro Naise",
		Email:    "saitoro@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$7iW14xiAFrOqlGcAvpoRBQ$H41kZ3e72u83FkyH9yFFvbqBKNFh8KyG7CfVaqZci24",
	},
	{
		Id:       4,
		Fullname: "Hercules Hernandes",
		Email:    "herce@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$RgNuO/ETgFjGwd8K76fEfw$xNp8UK+Kfl7wKIPr2/f0lcjYD9y4OLHMyvMF57o9U4c",
	},
	{
		Id:       5,
		Fullname: "Yunanda Ayunda",
		Email:    "yuyun@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$nbPIKPuVmz6j53HUoP1+3w$qVQRSmTbJkOiY6s/dCJFs19GQfrQVKnrspRp1RdsDeM",
	},
}

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "See All Users",
		Results: Listusers,
	})
}
func GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	row := Listusers

	var temp Users
	for _, data := range row {
		if data.Id == id {
			temp = data
		}
	}
	if temp.Id != id {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "ID Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "You Get User By ID",
		Results: temp,
	})

}
func LoginUsers(ctx *gin.Context) {
	var formUser Users
	ctx.ShouldBind(&formUser)

	var foundUser Users
	for _, getuser := range Listusers {
		if getuser.Email == formUser.Email {
			foundUser = getuser
		}
	}

	if foundUser == (Users{}) {
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

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for i, val := range Listusers {
		if val.Id == id {
			Listusers = append(Listusers[:i], Listusers[i+1:]...)
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Delete deleted successfully",
				Results: val,
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "ID Not Found",
	})

}

func EditUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var temp *Users
	for i := range Listusers {
		if Listusers[i].Id == id {
			temp = &Listusers[i]
		}
	}
	if temp == nil {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "ID Not Found",
		})
		return
	}
	var newData Users

	temp.Id = id

	if newData.Fullname != "" {
		temp.Fullname = newData.Fullname
	}
	if newData.Email != "" {
		temp.Email = newData.Email
	}
	if newData.Password != "" {
		temp.Password = newData.Password
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update Movie",
		Results: temp,
	})

}
func AddUser(ctx *gin.Context) {
	var formData Users
	ctx.ShouldBind(&formData)
	for _, getuser := range Listusers {
		if formData.Email == getuser.Email {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email Has Registered",
			})
		}
		if !strings.Contains(formData.Email, "@") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include @",
			})
			return
		}
		if len(formData.Password) < 6 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email Must be 6 Character",
			})
			return
		}
		if !strings.ContainsAny(formData.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include Uppercase Character",
			})
			return
		}
		if !strings.ContainsAny(formData.Password, "0123456789") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include One Number",
			})
			return
		}
		if !strings.ContainsAny(formData.Password, "!@#$%^&*()-_=+[]{}|;:,.<>?") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include Unique Character",
			})
			return
		}
	}
	hash, _ := argon2.CreateHash(formData.Password, formData.Password, argon2.DefaultParams)

	var row = Users{
		Fullname: formData.Fullname,
		Email:    formData.Email,
		Password: hash,
	}

	row.Id = len(Listusers) + 1
	Listusers = append(Listusers, row)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Your Movie Saved",
		Results: row,
	})
}
