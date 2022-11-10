package database

import (
	"database/sql"
	"fmt"
	"songs_api/models"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "arely"
	password = "secret"
	dbname   = "songs"
)

func SaveToDBJson(Data models.ResponseJson) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, _ := sql.Open("postgres", psqlconn)
	defer conn.Close()

	sqlAdd := `
		INSERT
		INTO "songs"(
			"id", 
			"name", 
			"artist", 
			"duration", 
			"album", 
			"artwork", 
			"price", 
			"origin"
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	for _, song := range Data.Songs {
		_, _ = conn.Exec(sqlAdd,
			song.ID,
			song.Name,
			song.Artist,
			song.Duration,
			song.Album,
			song.Artwork,
			song.Price,
			song.Origin,
		)
	}
}

func createLIKEQuery(query, column string) string {

	words := strings.Fields(query)
	var finalQuery string
	for _, word := range words {

		finalQuery += "UPPER( " + column + ") LIKE UPPER(" + " '%" + word + "%') OR "
		// fmt.Println(finalQuery)
	}
	return finalQuery
}

func ExecQuery(query string) []models.SongJson {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, _ := sql.Open("postgres", psqlconn)
	defer conn.Close()

	rows, _ := conn.Query(query)

	var result []models.SongJson
	for rows.Next() {
		item := models.SongJson{}
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.Artist,
			&item.Duration,
			&item.Album,
			&item.Artwork,
			&item.Price,
			&item.Origin,
		)
		result = append(result, item)
	}
	return result
}

func ReadFromDB(song, artist, album string) []models.SongJson {

	songLikeQuery := createLIKEQuery(song, "Name")
	artistLikeQuery := createLIKEQuery(artist, "Artist")
	albumLikeQuery := createLIKEQuery(album, "Album")
	likeQuery := songLikeQuery + artistLikeQuery + albumLikeQuery

	query := "SELECT * FROM songs WHERE " + likeQuery[:len(likeQuery)-3] + " LIMIT 10"

	response := ExecQuery(query)
	return response
}

func SaveToDBXML(Data models.ResponseXML) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, _ := sql.Open("postgres", psqlconn)
	defer conn.Close()

	sqlAdd := `
		INSERT
		INTO "songs"(
			"id", 
			"name", 
			"artist", 
			"duration", 
			"album", 
			"artwork", 
			"price", 
			"origin"
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	for _, song := range Data.Songs {
		_, _ = conn.Exec(sqlAdd,
			song.ID,
			song.Name,
			song.Artist,
			song.Duration,
			song.Album,
			song.Artwork,
			song.Price,
			song.Origin,
		)
	}

}

func InitDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, _ := sql.Open("postgres", psqlconn)
	defer conn.Close()

	// Se ejecuta una sola vez para crear la tabla
	statement, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS songs(
		ID TEXT,
		Name TEXT,
		Artist TEXT,
		Duration TEXT,
		Album TEXT,
		Artwork TEXT,
		Price TEXT,
		Origin TEXT
	)
	`)
	statement.Exec()

}
