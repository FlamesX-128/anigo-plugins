package main

import (
	"log"

	"github.com/FlamesX-128/anigo-plugins/monochinos/models/monoschinos"
	"github.com/FlamesX-128/anigo/models"
	"github.com/gocolly/colly"
)

type PackageModel struct{}

func (p PackageModel) Search(query string) []models.Anime {
	var data []models.Anime

	c := colly.NewCollector()

	c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
		data = append(data, models.Anime{
			Title: h.ChildText("h5.seristitles"),
			Id:    h.Attr("href"),
		})
	})

	if err := c.Visit(monoschinos.URL + "buscar?q=" + query); err != nil {
		log.Panicln(err)

	}

	return data
}

func (p PackageModel) Info(id string) models.Info {
	var data models.Info

	c := colly.NewCollector()

	c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
		data.Episodes = append(data.Episodes, models.Episode{
			Id: h.Attr("href"),
		})

	})

	if err := c.Visit(id); err != nil {
		log.Panicln(err)

	}

	return data
}

func (p PackageModel) Watch(id string) []models.Source {
	var data []models.Source

	return data
}

func (p PackageModel) Name() string {
	return "monochinos"
}

var Packages interface{} = []interface{}{
	PackageModel{},
}
