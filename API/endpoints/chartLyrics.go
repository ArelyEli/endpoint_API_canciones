package endpoints

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"songs_api/database"
	"songs_api/models"
	"strings"
)

func getSongs(artistName, songName string) models.ResponseXML {

	songNameChartLyrics := strings.Replace(songName, " ", "+", -1)
	artistNameChartLyrics := strings.Replace(artistName, " ", "+", -1)

	response, _ := http.Get("http://api.chartlyrics.com/apiv1.asmx/SearchLyric?artist=" + artistNameChartLyrics + "&song=" + songNameChartLyrics)

	data, _ := ioutil.ReadAll(response.Body)

	var DataObject models.ResponseXML
	xml.Unmarshal(data, &DataObject)
	// fmt.Println(DataObject)

	for i := 0; i < len(DataObject.Songs); i++ {
		DataObject.Songs[i].Origin = "chartLyrics"
		DataObject.Songs[i].Query = songName + " " + artistName
	}

	return DataObject
}
func GetChartLyricsData(artistName, songName string) {
	fmt.Println("Obteniendo canciones de ChartLyrics")
	data := getSongs(artistName, songName)
	database.SaveToDBXML(data)

}
