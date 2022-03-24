package mgnetmeutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/brandongallagher199/1337x-Bot-Go/internal/torrentserviceutils"
)

var linkShortenWaitGroup sync.WaitGroup

type MgnetmeResponse struct {
	State    string `json:"state"`
	Magnet   string `json:"magnet"`
	Shorturl string `json:"shorturl"`
	Message  string `json:"message"`
}

func GetMagnetLinks(torrentLinks []torrentserviceutils.TorrentServiceResponse) []torrentserviceutils.TorrentServiceResponse {
	linkShortenerChannel := make(chan string)
	for i := range torrentLinks {
		linkShortenWaitGroup.Add(1)
		go shortenLink(linkShortenerChannel, torrentLinks[i].Magnet)
	}

	go func() {
		linkShortenWaitGroup.Wait()
		close(linkShortenerChannel)
	}()

	var i int = 0
	for link := range linkShortenerChannel {
		torrentLinks[i].Magnet = link
		i++
	}

	return torrentLinks
}

func shortenLink(chnl chan string, magnetLink string) {
	response, err := http.Get("http://mgnet.me/api/create?&format=json&opt=&m=" + url.QueryEscape(magnetLink))
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject MgnetmeResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}
	chnl <- responseObject.Shorturl
	linkShortenWaitGroup.Done()
}
