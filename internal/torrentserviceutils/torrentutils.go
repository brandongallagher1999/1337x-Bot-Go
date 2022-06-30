package torrentserviceutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/brandongallagher199/1337x-Bot-Go/config"
)

var longMagnetWaitGroup sync.WaitGroup
var botConfig *config.Conf

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

type LongMagnetResponse struct {
	Magnet string `json:"magnet"`
}

type UpdateMagnetLink struct {
	Index int
	Link  string
}

func getLongMagnets(chnl chan UpdateMagnetLink, idx int, desc string) {
	response, err := http.Get("http://torrent-service:3000/longMagnet/" + url.QueryEscape(desc))
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject LongMagnetResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	chnl <- UpdateMagnetLink{Index: idx, Link: responseObject.Magnet}
	longMagnetWaitGroup.Done()
}

func QueryTorrentService(query string) []TorrentServiceResponse {
	longMagnetChannel := make(chan UpdateMagnetLink)
	response, err := http.Get("http://torrent-service:3000/" + url.QueryEscape(query))
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
	if len(responseObject) == 0 {
		return nil
	}

	for idx, torrent := range responseObject {
		longMagnetWaitGroup.Add(1)
		go getLongMagnets(longMagnetChannel, idx, torrent.Desc)
	}

	go func() {
		longMagnetWaitGroup.Wait()
		close(longMagnetChannel)
	}()

	for response := range longMagnetChannel {
		responseObject[response.Index].Magnet = response.Link
	}

	return responseObject
}
