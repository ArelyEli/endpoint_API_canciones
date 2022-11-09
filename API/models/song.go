package models

import "encoding/xml"

type ResponseJson struct {
	ResultCount int16      `json:"resultCount"`
	Songs       []SongJson `json:"results"`
}

type SongJson struct {
	ID       int32  `json:"trackId"`
	Name     string `json:"trackCensoredName"`
	Artist   string `json:"artistName"`
	Duration int32  `json:"trackTimeMillis"`
	Album    string `json:"collectionName"`
	Artwork  string `json:"collectionViewUrl"`
	Price    int16  `json:"trackPrice"`
	Origin   string
	Query    string
}

type ResponseXML struct {
	XMLName xml.Name  `xml:"ArrayOfSearchLyricResult" json:"-"`
	Songs   []SongXML `xml:"SearchLyricResult" json:"results"`
}

type SongXML struct {
	ID       int32  `xml:"TrackId" json:"TrackId"`
	Name     string `xml:"Song" json:"trackCensoredName"`
	Artist   string `xml:"Artist" json:"artistName"`
	Duration int32
	Album    string
	Artwork  string `xml:"SongUrl" json:"collectionViewUrl"`
	Price    int16
	Origin   string
	Query    string
}
