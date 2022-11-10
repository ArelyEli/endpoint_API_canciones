package database

import (
	"database/sql"
	"fmt"
	"songs_api/models"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func SaveToDBJsonL(Data models.ResponseJson) {

	const db string = "songs.db"

	conn, _ := sql.Open("sqlite3", db)

	// Se ejecuta una sola vez para crear la tabla
	statement, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS songs(
		ID VARCHAR(64),
		Name VARCHAR(64),
		Artist VARCHAR(64),
		Duration VARCHAR(64),
		Album VARCHAR(64),
		Artwork VARCHAR(64),
		Price VARCHAR(64),
		Origin VARCHAR(64)
	)
	`)
	statement.Exec()

	const sqlAdd = `INSERT INTO songs 
		(ID, Name, Artist, Duration, Album, Artwork, Price, Origin) 
		VALUES (?,?,?,?,?,?,?,?)`

	for _, song := range Data.Songs {
		// fmt.Println(song)
		statement, _ := conn.Prepare(sqlAdd)
		statement.Exec(
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

func createLIKEQueryL(query, column string) string {

	words := strings.Fields(query)
	var finalQuery string
	for _, word := range words {

		finalQuery += column + " LIKE" + " \"%" + word + "%\" OR "
		// fmt.Println(finalQuery)
	}
	return finalQuery
}

func ExecQueryL(query string) []models.SongJson {
	const db string = "songs.db"

	conn, _ := sql.Open("sqlite3", db)

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
	fmt.Println(result)
	return result
}

func ReadFromDBL(song, artist, album string) []models.SongJson {

	songLikeQuery := createLIKEQuery(song, "Name")
	artistLikeQuery := createLIKEQuery(artist, "Artist")
	albumLikeQuery := createLIKEQuery(album, "Album")
	likeQuery := songLikeQuery + artistLikeQuery + albumLikeQuery

	query := "SELECT * FROM songs WHERE " + likeQuery[:len(likeQuery)-3] + " LIMIT 10"

	response := ExecQuery(query)
	return response
}

func SaveToDBXMLL(Data models.ResponseXML) {

	const db string = "songs.db"

	conn, _ := sql.Open("sqlite3", db)

	// Se ejecuta una sola vez para crear la tabla
	statement, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS songs(
		ID VARCHAR(64),
		Name VARCHAR(64),
		Artist VARCHAR(64),
		Duration VARCHAR(64),
		Album VARCHAR(64),
		Artwork VARCHAR(64),
		Price VARCHAR(64),
		Origin VARCHAR(64)
	)
	`)
	statement.Exec()

	const sqlAdd = `INSERT INTO songs 
		(ID, Name, Artist, Duration, Album, Artwork, Price, Origin) 
		VALUES (?,?,?,?,?,?,?,?)`

	for _, song := range Data.Songs {
		// fmt.Println(song)
		statement, _ := conn.Prepare(sqlAdd)
		statement.Exec(
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

func InitDBL() {
	const db string = "songs.db"

	conn, _ := sql.Open("sqlite3", db)

	// Se ejecuta una sola vez para crear la tabla
	statement, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS songs(
		ID VARCHAR(64),
		Name VARCHAR(64),
		Artist VARCHAR(64),
		Duration VARCHAR(64),
		Album VARCHAR(64),
		Artwork VARCHAR(64),
		Price VARCHAR(64),
		Origin VARCHAR(64)
	)
	`)
	statement.Exec()

}
