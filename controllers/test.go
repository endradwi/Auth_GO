package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type Movie struct {
// 	Tittle      string `json:"tittle"`
// 	Description string `json:"description"`
// 	Image       string `json:"image"`
// }

// type Response struct {
// 	Success bool    `json:"success"`
// 	Message string  `json:"message"`
// 	Results []Movie `json:"results"`
// }

// var ListView = []Movie{
// 	// Sample movie data
// 	{Tittle: "Movie 1", Description: "Description 1", Image: "Image 1"},
// 	{Tittle: "Movie 2", Description: "Description 2", Image: "Image 2"},
// 	// Add more movies as needed
// }

// // Function to filter movies based on search parameters
// func searchMovies(movies []Movie, title, desc, image string) []Movie {
// 	var filteredMovies []Movie
// 	for _, movie := range movies {
// 		if (title == "" || title == movie.Tittle) &&
// 			(desc == "" || desc == movie.Description) &&
// 			(image == "" || image == movie.Image) {
// 			filteredMovies = append(filteredMovies, movie)
// 		}
// 	}
// 	return filteredMovies
// }

// // Function to handle pagination
// func paginateMovies(movies []Movie, page, limit int) []Movie {
// 	startIndex := (page - 1) * limit
// 	endIndex := startIndex + limit

// 	// Check if the start index is within bounds
// 	if startIndex > len(movies) {
// 		return []Movie{} // No movies to return if out of range
// 	}

// 	// Ensure the end index doesn't exceed the list length
// 	if endIndex > len(movies) {
// 		endIndex = len(movies)
// 	}

// 	return movies[startIndex:endIndex]
// }

// // Function to handle the sorting of movies (if applicable)
// func sortMovies(movies []Movie) []Movie {
// 	// Implement your sorting logic here (e.g., by title, release date, etc.)
// 	// In this case, we'll just return the original order
// 	return movies
// }

// func GetAllMovies(ctx *gin.Context) {
// 	// Get query parameters
// 	searchTittle := ctx.DefaultQuery("tittle", "")
// 	searchDesc := ctx.DefaultQuery("desc", "")
// 	searchImage := ctx.DefaultQuery("image", "")
// 	pageStr := ctx.DefaultQuery("page", "1")
// 	limitStr := ctx.DefaultQuery("limit", "10")

// 	// Convert page and limit to integers
// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil || page <= 0 {
// 		page = 1
// 	}

// 	limit, err := strconv.Atoi(limitStr)
// 	if err != nil || limit <= 0 {
// 		limit = 10
// 	}

// 	// Step 1: Filter movies based on search parameters
// 	filteredMovies := searchMovies(ListView, searchTittle, searchDesc, searchImage)

// 	// Step 2: Sort movies (if needed)
// 	sortedMovies := sortMovies(filteredMovies)

// 	// Step 3: Paginate the sorted, filtered movies
// 	paginatedMovies := paginateMovies(sortedMovies, page, limit)

// 	// Return the result as JSON
// 	if len(paginatedMovies) == 0 {
// 		ctx.JSON(http.StatusOK, Response{
// 			Success: true,
// 			Message: "No Movies Found",
// 			Results: []Movie{},
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, Response{
// 		Success: true,
// 		Message: "Movies Retrieved Successfully",
// 		Results: paginatedMovies,
// 	})
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/movies", GetAllMovies)
// 	r.Run(":8080")
// }
