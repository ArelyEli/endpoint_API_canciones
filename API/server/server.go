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

	artistNameLower := strings.ToLower(artistName)
	songNameLower := strings.ToLower(songName)

	token := c.Query("token")

	if artistName == "" || songName == "" || token == "" {
		c.String(http.StatusBadRequest, "artist, song and token are required")
	} else if IsAValidateToken(token) {
		data := database.Read_from_db(songNameLower + " " + artistNameLower)
		if data == nil {
			endpoints.Get_itunes_Data(songNameLower + " " + artistNameLower)
			endpoints.GetChartLyricsData(artistNameLower, songNameLower)
			data := database.Read_from_db(songNameLower + " " + artistNameLower)
			c.IndentedJSON(http.StatusOK, data)
		} else {
			c.IndentedJSON(http.StatusOK, data)
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
