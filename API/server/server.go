package server

import (
	"fmt"
	"net/http"
	"songs_api/database"
	"songs_api/endpoints"
	"strings"

	"github.com/gin-gonic/gin"
)

func getCanciones(c *gin.Context) {
	artistName := c.Query("artist")
	songName := c.Query("song")
	albumName := c.Query("album")

	artistNameLower := strings.ToLower(artistName)
	songNameLower := strings.ToLower(songName)
	albumNameLower := strings.ToLower(albumName)

	token := c.Query("token")

	if token == "" {
		c.String(http.StatusBadRequest, "token is required")

	} else if IsAValidateToken(token) {

		if artistNameLower == "" && songNameLower == "" && albumNameLower == "" {
			query := "SELECT * FROM songs LIMIT 10"
			response := database.ExecQuery(query)
			c.IndentedJSON(http.StatusOK, response)
		} else {

			data := database.ReadFromDB(songNameLower, artistNameLower, albumNameLower)
			if data == nil {
				fmt.Println("New Search")
				endpoints.Get_itunes_Data(songNameLower + " " + artistNameLower + " " + albumNameLower)
				endpoints.GetChartLyricsData(artistNameLower, songNameLower)
				data := database.ReadFromDB(songNameLower, artistNameLower, albumNameLower)
				c.IndentedJSON(http.StatusOK, data)
			} else {
				c.IndentedJSON(http.StatusOK, data)
			}
		}
	} else {
		c.String(http.StatusBadRequest, "token is invalide")
	}
}

func singup(c *gin.Context) {
	fmt.Print("Creating new account...")
	username := c.Query("username")
	if username == "" {
		c.String(http.StatusBadRequest, "username is required")
	} else {
		tokenString := GenerateToken(username)
		c.String(http.StatusOK, tokenString)
	}

}

func Server() {
	fmt.Println("iniciando servidor...")
	database.InitDB()

	router := gin.Default()
	router.GET("/canciones", getCanciones)
	router.POST("/singup", singup)
	router.Run("0.0.0.0:6767")
}
