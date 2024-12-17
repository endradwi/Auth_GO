package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results ,omitempty"`
}

type Movie struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle" form:"tittle"`
	Image       string `json:"image" form:"images"`
	Description string `json:"desc" form:"desc"`
}

var ListMovie []Movie = []Movie{
	{
		Id:          1,
		Tittle:      "Spiderman",
		Image:       "http://urlspiderman",
		Description: "Lorem ipsum no",
	},
	{
		Id:          2,
		Tittle:      "Black Widow",
		Image:       "http://urlblackwidow",
		Description: "Lorem ipsum nor found",
	},
	{
		Id:          3,
		Tittle:      "Spiderman",
		Image:       "http://urlspiderman",
		Description: "Lorem ipsum nor found",
	},
	{
		Id:          4,
		Tittle:      "Black Widow",
		Image:       "http://urlblackwidow",
		Description: "Lorem ipsum nor found",
	},
	{
		Id:          5,
		Tittle:      "Spiderman",
		Image:       "http://urlspiderman",
		Description: "Lorem ipsum nor found",
	},
	{
		Id:          6,
		Tittle:      "Black Widow",
		Image:       "http://urlspider",
		Description: "Lorem ipsum nor found",
	},
}

func GetAllMovies(ctx *gin.Context) {
	search := ctx.Query("search")
	searchpage := ctx.Query("page")
	searchlimit := ctx.Query("limit")
	sortmovie := ctx.Query("sort")

	pagesearch, _ := strconv.Atoi(searchpage)
	limitsearch, _ := strconv.Atoi(searchlimit)
	if pagesearch <= 0 {
		pagesearch = 1
	}
	if limitsearch <= 0 {
		limitsearch = 10
	}

	offset := (pagesearch - 1) * limitsearch
	end := offset + limitsearch

	if end > len(ListMovie) {
		end = len(ListMovie)
	}
	if sortmovie != "" {
		sort.Slice(ListMovie, func(i, j int) bool {
			if sortmovie == "asc" {
				return ListMovie[i].Tittle < ListMovie[j].Tittle
			} else if sortmovie == "desc" {
				return ListMovie[i].Tittle > ListMovie[j].Tittle
			}
			return ListMovie[i].Tittle < ListMovie[j].Tittle
		})
	}
	var row []Movie

	for _, data := range ListMovie[offset:end] {
		if search == "" ||
			strings.Contains(strings.ToLower(data.Tittle), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(data.Image), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(data.Description), strings.ToLower(search)) {
			row = append(row, data)
		}
	}

	ctx.JSON(200, Response{
		Success: true,
		Message: "See All Movies",
		Results: row,
	})
}

func EditMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var temp *Movie
	for i := range ListMovie {
		if ListMovie[i].Id == id {
			temp = &ListMovie[i]
		}
	}
	if temp == nil {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "ID Not Found",
		})
		return
	}
	var newData Movie

	temp.Id = id

	if newData.Tittle != "" {
		temp.Tittle = newData.Tittle
	}
	if newData.Image != "" {
		temp.Image = newData.Image
	}
	if newData.Description != "" {
		temp.Description = newData.Description
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update Movie",
		Results: temp,
	})

}

func GetMoviesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	row := ListMovie

	var temp Movie
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
		Message: "You Get Movie By ID",
		Results: temp,
	})

}

func SaveMovies(ctx *gin.Context) {
	var formData Movie
	ctx.ShouldBind(&formData)

	var row = Movie{
		Tittle:      formData.Tittle,
		Image:       formData.Image,
		Description: formData.Description,
	}

	row.Id = len(ListMovie) + 1
	ListMovie = append(ListMovie, row)
	ctx.JSON(200, Response{
		Success: true,
		Message: "Your Movie Saved",
		Results: row,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for i, val := range ListMovie {
		if val.Id == id {
			ListMovie = append(ListMovie[:i], ListMovie[i+1:]...)
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
