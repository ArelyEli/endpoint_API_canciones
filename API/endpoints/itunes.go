package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"songs_api/database"
	"songs_api/models"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func get_songs(query string) models.ResponseJson {
	fmt.Println("Obteniendo cancionces de iTunes")

	query_to_itunes := strings.Replace(query, " ", "+", -1)

	response, err := http.Get("https://itunes.apple.com/search?term=" + query_to_itunes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var DataObject models.ResponseJson
	json.Unmarshal(data, &DataObject)

	// Modificación de valores
	for i := 0; i < len(DataObject.Songs); i++ {
		DataObject.Songs[i].Origin = "Apple"
		DataObject.Songs[i].Query = query
	}
	return DataObject
}

func Get_itunes_Data(query string) {
	Data := get_songs(query)

	database.SaveToDBJson(Data)
}
