package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func FetchToByte(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("http.Get(url) error : %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)

	return ioutil.ReadAll(bodyReader)
}

func Fetch(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}
	//bodyReader := bufio.NewReader(resp.Body)

	return goquery.NewDocumentFromResponse(resp)
}
func FetchByForm(targetUrl string, values url.Values) (*goquery.Document, error) {

	resp, err := http.PostForm(targetUrl, values)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return goquery.NewDocumentFromResponse(resp)
}
