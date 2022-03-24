package mgnetmeutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var linkShortenWaitGroup sync.WaitGroup

type MgnetmeResponse struct {
	State    string `json:"state"`
	Magnet   string `json:"magnet"`
	Shorturl string `json:"shorturl"`
	Message  string `json:"message"`
}

func GetMagnetLinks(magnetLinks []string) []string {
	linkShortenerChannel := make(chan string)
	for i := range magnetLinks {
		linkShortenWaitGroup.Add(1)
		go shortenLink(linkShortenerChannel, magnetLinks[i])
	}

	go func() {
		linkShortenWaitGroup.Wait()
		close(linkShortenerChannel)
	}()
	shortenedLinks := make([]string, 0)
	for link := range linkShortenerChannel {
		shortenedLinks = append(shortenedLinks, link)
	}

	return shortenedLinks
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
