package torrentserviceutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type TorrentSeriveResponse struct {
	State string `json:"state"`
}

func QueryTorrentService(query string) {
	response, err := http.Get("" + url.QueryEscape(query))
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject TorrentSeriveResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}
}
