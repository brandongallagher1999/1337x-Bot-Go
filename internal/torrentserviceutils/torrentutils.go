package torrentserviceutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var longMagnetWaitGroup sync.WaitGroup

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

func getLongMagnets(chnl chan string, desc string) {
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

	chnl <- responseObject.Magnet
	longMagnetWaitGroup.Done()
}

func QueryTorrentService(query string) []TorrentServiceResponse {
	longMagnetChannel := make(chan string)
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

	for i := range responseObject {
		longMagnetWaitGroup.Add(1)
		go getLongMagnets(longMagnetChannel, responseObject[i].Desc)
	}

	go func() {
		longMagnetWaitGroup.Wait()
		close(longMagnetChannel)
	}()

	var j int = 0
	for mgnt := range longMagnetChannel {
		responseObject[j].Magnet = mgnt
		j++
	}

	return responseObject
}
