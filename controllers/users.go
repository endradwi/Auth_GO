package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"test/models"

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
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "2"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	sortuser := ctx.DefaultQuery("sort", "ASC")
	searchall := models.FindAllUser(page, limit, search, sortuser)

	if sortuser != "ASC" {
		sortuser = "DESC"
	}
	count := models.CountData(search)

	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	prevPage := page - 1
	if prevPage < 2 {
		prevPage = 0
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: "See All Users",
		PageInfo: PageInfo{
			CurentPage: page,
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPage:  totalPage,
			TotalData:  count,
		},
		Results: searchall,
	})

	// totalData :=
	// pagesearch, _ := strconv.Atoi(searchpage)
	// limitsearch, _ := strconv.Atoi(searchlimit)
	// if pagesearch <= 0 {
	// 	pagesearch = 1
	// }
	// if limitsearch <= 0 {
	// 	limitsearch = 10
	// }

	// offset := (pagesearch - 1) * limitsearch
	// end := offset + limitsearch

	// if end > len(Listusers) {
	// 	end = len(Listusers)
	// }
	// if sortmovie != "" {
	// 	sort.Slice(Listusers, func(i, j int) bool {
	// 		if sortmovie == "asc" {
	// 			return Listusers[i].Fullname < Listusers[j].Fullname
	// 		} else if sortmovie == "desc" {
	// 			return Listusers[i].Fullname > Listusers[j].Fullname
	// 		}
	// 		return Listusers[i].Fullname < Listusers[j].Fullname
	// 	})
	// }
	// var row []Users

	// for _, data := range Listusers[offset:end] {
	// 	if search == "" ||
	// 		strings.Contains(strings.ToLower(data.Fullname), strings.ToLower(search)) ||
	// 		strings.Contains(strings.ToLower(data.Email), strings.ToLower(search)) ||
	// 		strings.Contains(strings.ToLower(data.Password), strings.ToLower(search)) {
	// 		row = append(row, data)
	// 	}
	// }

}
func GetUserById(ctx *gin.Context) {
	pram, _ := strconv.Atoi(ctx.Param("id"))
	param := models.FindOneUser(pram)
	// row := Listusers
	// var temp Users
	// for _, data := range err {
	// 	if data.Id == id {
	// 		temp = data
	// 	}
	// }
	// if temp.Id != id {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "You Get User By ID",
		Results: param,
	})

}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	params := models.FindOneUser(id)

	if params == (models.Users{}) {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: true,
			Message: "User Not Found",
			Results: params,
		})
		return
	}

	deleted := models.DeleteUser(id)
	// for i, val := range Listusers {
	// 	if val.Id == id {
	// 		Listusers = append(Listusers[:i], Listusers[i+1:]...)
	// 		ctx.JSON(http.StatusOK, Response{
	// 			Success: true,
	// 			Message: "Delete deleted successfully",
	// 			Results: val,
	// 		})
	// 		return
	// 	}
	// }
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Delete Success",
		Results: deleted,
	})

}

func EditUser(ctx *gin.Context) {
	paramId, _ := strconv.Atoi(ctx.Param("id"))
	user := models.FindOneUser(paramId)
	fmt.Println(user)
	if user == (models.Users{}) {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "ID Not Found",
		})
		return
	}
	ctx.ShouldBind(&user)

	if !strings.Contains(user.Password, "$argon2i$v=19") {
		hash, _ := argon2.CreateHash(user.Password, user.Password, argon2.DefaultParams)
		if user.Password != "" {
			user.Password = hash
		}
	}
	updated := models.UpdateUser(user)
	fmt.Println(updated)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: updated,
	})
	// var temp *Users
	// for i := range Listusers {
	// 	if Listusers[i].Id == id {
	// 		temp = &Listusers[i]
	// 	}
	// }
	// if temp == nil {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// 	})
	// 	return
	// }
	// var newData Users

	// temp.Id = id

	// if newData.Fullname != "" {
	// 	temp.Fullname = newData.Fullname
	// }
	// if newData.Email != "" {
	// 	temp.Email = newData.Email
	// }
	// if newData.Password != "" {
	// 	temp.Password = newData.Password
	// }

}
func AddUser(ctx *gin.Context) {
	var formData models.Users
	ctx.ShouldBind(&formData)
	fmt.Println(formData)
	hash, _ := argon2.CreateHash(formData.Password, formData.Password, argon2.DefaultParams)
	if formData.Password != "" {
		formData.Password = hash
	}

	new := models.InsertUser(formData)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Your User Saved",
		Results: new,
	})
	// var row = Users{
	// 	Fullname: formData.Fullname,
	// 	Email:    formData.Email,
	// 	Password: hash,
	// }
	// for _, getuser := range Listusers {
	// 	if formData.Email == getuser.Email {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Email Has Registered",
	// 		})
	// 	}
	// 	if !strings.Contains(formData.Email, "@") {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Password Must Include @",
	// 		})
	// 		return
	// 	}
	// 	if len(formData.Password) < 6 {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Email Must be 6 Character",
	// 		})
	// 		return
	// 	}
	// 	if !strings.ContainsAny(formData.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Password Must Include Uppercase Character",
	// 		})
	// 		return&delete.Id
	// 	}
	// 	if !strings.ContainsAny(formData.Password, "0123456789") {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Password Must Include One Number",
	// 		})
	// 		return
	// 	}
	// 	if !strings.ContainsAny(formData.Password, "!@#$%^&*()-_=+[]{}|;:,.<>?") {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Message: "Password Must Include Unique Character",
	// 		})
	// 		return
	// 	}
	// }

	// row.Id = len(Listusers) + 1
	// Listusers = append(Listusers, row)
	// ctx.JSON(http.StatusOK, Response{
	// 	Success: true,
	// 	Message: "Your User Saved",
	// 	Results: new,
	// })
}
