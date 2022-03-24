package torrentserviceutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type TorrentServiceResponse struct {
	Title    string `json:"title"`
	Time     string `json:"time"`
	Seeds    int    `json:"seeds"`
	Peers    int    `json:"peers"`
	Size     string `json:"size"`
	Desc     string `json:"desc"`
	Provider string `json:"provider"`
	Magnet   string `json:"magnet"`
	Number   int    `json:"number"`
}

func QueryTorrentService(query string) []TorrentServiceResponse {
	response, err := http.Get("http://localhost:3000/" + url.QueryEscape(query))
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []TorrentServiceResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}
	return responseObject
}
