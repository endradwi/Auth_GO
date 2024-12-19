package models

import (
	"context"
	"fmt"
	"test/lib"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movies struct {
	Id           int       `json:"id"`
	Tittle       string    `json:"tittle" form:"tittle"`
	Genre        string    `json:"genre" form:"genre"`
	Synopsis     string    `json:"synopsis" form:"synopsis"`
	Author       string    `json:"author" form:"author"`
	Actors       string    `json:"actors" form:"actors"`
	Release_date time.Time `json:"release_date" form:"release_date"`
	Duration     time.Time `json:"duration" form:"duration"`
	Created_at   time.Time `json:"created_at" form:"created_at"`
	Image        string    `json:"image" form:"image"`
	// Updated_at   time.Time `json:"updated_at"`
}

type ListMovie []Movies

func FindOneMovie(paramId int) Movies {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var movie Movies

	conn.QueryRow(context.Background(), `
SELECT id, tittle, genre, synopsis, author,
actors, release_date, duration, created_at
FROM movies WHERE id = $1
`, paramId).Scan(&movie.Id, &movie.Tittle, &movie.Genre,
		&movie.Synopsis, &movie.Author, &movie.Actors,
		&movie.Release_date, &movie.Duration, &movie.Created_at)
	return movie
}

func CountMovie(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int
	search = fmt.Sprintf("%%%s%%", search)

	temp := conn.QueryRow(context.Background(), `
SELECT COUNT(id)
FROM movies
WHERE tittle ILIKE $1
`, search).Scan((&count))
	fmt.Println(temp)
	return count
}
func FindAllMovie(page int, limit int, search string, sort string) ListMovie {
	conn := lib.DB()
	defer conn.Close(context.Background())
	offset := (page - 1) * limit

	searching := fmt.Sprintf("%%%s%%", search)

	query := fmt.Sprintf(`
	SELECT id, tittle, genre, synopsis,
	author, actors, release_date,
	duration, created_at, '' as image
	FROM movies
	WHERE tittle ILIKE $1
 	ORDER BY id %s
 	LIMIT $2 OFFSET $3
 `, sort)
	rows, _ := conn.Query(context.Background(), query, searching, limit, offset)
	movie, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Movies])
	return movie
}

func UpdateMovie(movie Movies) Movies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var update Movies

	conn.QueryRow(context.Background(), `
 UPDATE movies SET tittle=$1, genre=$2, synopsis=$3,
 author=$4, actors=$5,release_date=&6, duration=&7 WHERE id = $8

 RETURNING id, tittle, genre, synopsis, author,
 actors, release_date, duration, created_at
 `, movie.Tittle, movie.Genre, movie.Synopsis, movie.Author, movie.Actors,
		movie.Release_date, movie.Duration, movie.Id).Scan(
		&update.Id,
		&update.Tittle,
		&update.Genre,
		&update.Synopsis,
		&update.Author,
		&update.Actors,
		&update.Release_date,
		&update.Duration,
	)
	fmt.Println(update)
	return update
}

func DeleteMovie(id int) Movies {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var delete Movies
	conn.QueryRow(context.Background(), `
DELETE FROM movies WHERE id=$1
RETURNING id, tittle, genre, synopsis,
author, actors,release_date, duration, created_at, updated_at
`, id).Scan(
		&delete.Id,
		&delete.Tittle,
		&delete.Genre,
		&delete.Synopsis,
		&delete.Author,
		&delete.Actors,
		&delete.Release_date,
		&delete.Duration,
	)
	return delete
}
func InsertMovie(movie Movies) *Movies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new Movies

	conn.QueryRow(context.Background(), `
        INSERT INTO movies (tittle, genre, synopsis, author, actors,
		release_date, duration, created_at, image) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, tittle, genre, synopsis, author, actors,
		release_date, duration, created_at, image
    `, movie.Tittle, movie.Genre, movie.Synopsis, movie.Author, movie.Actors,
		movie.Release_date, movie.Duration, movie.Created_at, movie.Image).Scan(
		&new.Id, &new.Tittle, &new.Genre, &new.Synopsis, &new.Author,
		&new.Actors, &new.Release_date, &new.Duration, &new.Created_at, &new.Image)

	return &new
}
