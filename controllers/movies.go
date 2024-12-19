package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
)

type PageInfo struct {
	CurentPage int `json:"current_page"`
	NextPage   int `json:"next_page"`
	PrevPage   int `json:"prev_page"`
	TotalPage  int `json:"total_page"`
	TotalData  int `json:"total_data"`
}

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"pageinfo ,omitempty"`
	Results  any    `json:"results ,omitempty"`
}

// type Movie struct {
// 	Id          int    `json:"id"`
// 	Tittle      string `json:"tittle" form:"tittle"`
// 	Image       string `json:"image" form:"images"`
// 	Description string `json:"desc" form:"desc"`
// }

// var ListMovie []Movie

func GetAllMovies(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	sortmovie := ctx.DefaultQuery("sort", "ASC")
	if sortmovie != "ASC" {
		sortmovie = "DESC"
	}

	var movies models.ListMovie
	var count int
	get := lib.Redis().Get(context.Background(), ctx.Request.RequestURI)
	getCount := lib.Redis().Get(context.Background(),
		fmt.Sprintf("count+%s", ctx.Request.RequestURI))
	if get.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &movies)
	} else {
		movies = models.FindAllMovie(page, limit, search, sortmovie)
		change, _ := json.Marshal(movies)
		lib.Redis().Set(
			context.Background(),
			ctx.Request.RequestURI,
			string(change),
			0,
		)
	}
	if getCount.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &count)
	} else {
		count = models.CountData(search)
		change, _ := json.Marshal(count)
		lib.Redis().Set(context.Background(),
			fmt.Sprintf("count+%s", ctx.Request.RequestURI),
			string(change),
			0,
		)
	}
	// count = models.CountData(search)
	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	log.Println("errorrrr", totalPage)
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
		Results: movies,
	})
}

// func EditMovie(ctx *gin.Context) {
// id, _ := strconv.Atoi(ctx.Param("id"))

// var temp *Movie
// for i := range ListMovie {
// if ListMovie[i].Id == id {
// temp = &ListMovie[i]
// }
// }
// if temp == nil {
// ctx.JSON(http.StatusNotFound, Response{
// Success: false,
// Message: "ID Not Found",
// })
// return
// }
// var newData Movie
//
// temp.Id = id
//
// if newData.Tittle != "" {
// temp.Tittle = newData.Tittle
// }
// if newData.Image != "" {
// temp.Image = newData.Image
// }
// if newData.Description != "" {
// temp.Description = newData.Description
// }
//
// ctx.JSON(http.StatusOK, Response{
// Success: true,
// Message: "Update Movie",
// Results: temp,
// })
//
// }

func GetMoviesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	find := models.FindOneMovie(id)
	// row := ListMovie

	// var temp Movie
	// for _, data := range row {
	// 	if data.Id == id {
	// 		temp = data
	// 	}
	// }
	// if temp.Id != id {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// }

	// row.Id = len(ListMovie) + 1
	// ListMovie = append(ListMovie, row)
	// ctx.JSON(200, Response{
	// Success: true,
	// Message: "Your Movie Saved",
	// Results: temp,
	// })
	// }
	//
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "You Get Movie By ID",
		Results: find,
	})

}

func SaveMovies(ctx *gin.Context) {
	var formData models.Movies
	ctx.ShouldBind(&formData)
	// file, _ := ctx.FormFile("image")
	// filename := uuid.New().String()
	// splitfile := strings.Split(file.Filename, ".")[1]
	// ext := splitfile[len(splitfile)-1]
	// storedFile := fmt.Sprintf("%s.%s", filename, ext)
	// if file.Size > 2<<8 {
	// ctx.JSON(400, Response{
	// Success: false,
	// Message: "Image to large",
	// })
	// return
	// }
	// ctx.SaveUploadedFile(file, fmt.Sprintf("uploads/movies/%s", storedFile))
	// if file.Filename != "" {
	// formData.Image = storedFile
	// }

	temp := models.InsertMovie(formData)

	data := lib.Redis().Scan(context.Background(), 0, "", 0).Iterator()
	for data.Next(context.Background()) {

	}
	// var row = Movie{
	// 	Tittle:      formData.Tittle,
	// 	Image:       formData.Image,
	// 	Description: formData.Description,
	// }

	// row.Id = len(ListMovie) + 1
	// ListMovie = append(ListMovie, row)
	ctx.JSON(200, Response{
		Success: true,
		Message: "Your Movie Saved",
		Results: temp,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	deleted := models.DeleteMovie(id)
	//	for i, val := range ListMovie {
	//		if val.Id == id {
	//			ListMovie = append(ListMovie[:i], ListMovie[i+1:]...)
	//			ctx.JSON(http.StatusOK, Response{
	//				Success: true,
	//				Message: "Delete deleted successfully",
	//				Results: val,
	//			})
	//			return
	//		}
	//	}
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Deleted Success",
		Results: deleted,
	})

}
