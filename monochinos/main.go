package main

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo-plugins/monochinos/controllers"
	"github.com/FlamesX-128/anigo-plugins/monochinos/controllers/providers"
	"github.com/FlamesX-128/anigo-plugins/monochinos/models/monoschinos"
	"github.com/PuerkitoBio/goquery"
)

type PackageModel struct{}

func (p PackageModel) Search(query string) []models.Anime {
	var url string = monoschinos.URL + "buscar?q=" + query
	var target string = "div.row a[href]"

	data, _ := controllers.Scrape(url, target, func(i int, s *goquery.Selection) models.Anime {
		a := models.Anime{
			Title: s.Find("h3.seristitles").Text(),
			Id:    s.AttrOr("href", ""),
		}

		return a
	})

	return data
}

func (p PackageModel) Info(id string) models.Info {
	var target string = "div.allanimes"

	data, _ := controllers.Scrape(id, target, func(i int, s *goquery.Selection) (data models.Info) {
		var target string = "div.col-item"

		s.Find(target).Each(func(i int, s *goquery.Selection) {
			number, _ := strconv.Atoi(s.AttrOr("data-episode", "0"))

			data.Episodes = append(data.Episodes, models.Episode{
				Number: uint32(number),
				Id:     s.Find("a[href]").AttrOr("href", ""),
			})
		})

		return
	})

	return data[0]
}

func (p PackageModel) Watch(id string) []models.Source {
	var target string = "div.heromain ul.dropcaps li#play-video a[data-player]"

	data, _ := controllers.Scrape(id, target, func(i int, s *goquery.Selection) models.Source {
		quality, _ := base64.StdEncoding.DecodeString(s.AttrOr("data-player", ""))
		params := strings.Split(string(quality), "?url=")

		return models.Source{
			Quality: s.Text(),
			Url:     params[(len(params) - 1)],
		}
	})

	data, _ = providers.ProviderHandler(data)

	return data
}

var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"monochinos": PackageModel{},
	},
}
