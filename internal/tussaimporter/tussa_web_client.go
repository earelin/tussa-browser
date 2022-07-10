package tussaimporter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const tussaBaseUrl = "https://tussa.org/web/"

func downloadAllLinesPage() string {
	allLinesPageUrl := fmt.Sprintf("%stransporte.php?l=todas", tussaBaseUrl)
	return extractUrlSourceCode(allLinesPageUrl)
}

func downloadLinePage(id int) string {
	linePageUrl := fmt.Sprintf("%slinas.php?id=%d", tussaBaseUrl, id)
	return extractUrlSourceCode(linePageUrl)
}

func downloadStopsPage() string {
	stopsPageUrl := fmt.Sprintf("%stransporte.php", tussaBaseUrl)
	return extractUrlSourceCode(stopsPageUrl)
}

func extractUrlSourceCode(url string) string {
	response := getUrlResponse(url)
	return getResponseBodyString(response)
}

func getUrlResponse(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Cannot download from %s: %s", url, err))
	}

	return response
}

func getResponseBodyString(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("Cannot read body from request: %s", err))
	}

	defer response.Body.Close()

	return string(body)
}
