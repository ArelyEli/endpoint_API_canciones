package database

import (
	"database/sql"
	"songs_api/models"

	_ "github.com/mattn/go-sqlite3"
)

func SaveToDBJson(Data models.ResponseJson) {

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
		Origin VARCHAR(64),
		Query VARCHAR(64)
	)
	`)
	statement.Exec()

	const sqlAdd = `INSERT INTO songs 
		(ID, Name, Artist, Duration, Album, Artwork, Price, Origin, Query) 
		VALUES (?,?,?,?,?,?,?,?,?)`

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
			song.Query)
	}

}

func Read_from_db(query string) []models.SongJson {
	const db string = "songs.db"

	conn, _ := sql.Open("sqlite3", db)

	rows, _ := conn.Query(`
		SELECT * 
		FROM songs
		WHERE Query = ?
	`, query)

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
			&item.Query,
		)
		result = append(result, item)
	}
	// fmt.Println(result)
	return result

}

func SaveToDBXML(Data models.ResponseXML) {

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
		Origin VARCHAR(64),
		Query VARCHAR(64)
	)
	`)
	statement.Exec()

	const sqlAdd = `INSERT INTO songs 
		(ID, Name, Artist, Duration, Album, Artwork, Price, Origin, Query) 
		VALUES (?,?,?,?,?,?,?,?,?)`

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
			song.Query)
	}

}

func InitDB() {
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
		Origin VARCHAR(64),
		Query VARCHAR(64)
	)
	`)
	statement.Exec()

}