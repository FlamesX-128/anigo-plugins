package controllers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scrape[T interface{}](url string, target string, callback func(i int, s *goquery.Selection) T) (data []T, err error) {
	var doc *goquery.Document
	var res *http.Response

	if res, err = http.Get(url); err != nil {

		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(
			"failed to fetch data from %s, status code: %d", url, res.StatusCode,
		)
	}

	if doc, err = goquery.NewDocumentFromReader(res.Body); err != nil {

		return
	}

	doc.Find(target).Each(func(i int, s *goquery.Selection) {
		data = append(data, callback(i, s))
	})

	return
}
