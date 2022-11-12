package gogoanime

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/FlamesX-128/anigo-plugins/consumet/models/gogoanime"
	"github.com/FlamesX-128/anigo-plugins/models"
)

type PackageModel struct{}

func NewJSONRequest(method string, url string) *http.Response {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panicln(err)
	}

	return resp
}

func Query[T gogoanime.Response](route string) T {
	resp := NewJSONRequest("GET", (gogoanime.URL + route))
	defer resp.Body.Close()

	var data T

	bodyBytes, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		log.Panicln(err)
	}

	return data
}

func (p PackageModel) Search(query string) []models.Anime {
	var results []models.Anime

	for _, data := range Query[gogoanime.Search](query).Results {
		results = append(results, models.Anime{
			Title: data.Title,
			Id:    data.Id,
		})
	}

	return results
}

func (p PackageModel) Info(id string) models.Info {
	var info models.Info

	for _, data := range Query[gogoanime.Info]("info/" + id).Episodes {
		info.Episodes = append(info.Episodes, models.Episode{
			Number: data.Number,
			Id:     data.Id,
		})
	}

	return info
}

func (p PackageModel) Watch(id string) []models.Source {
	var source []models.Source

	for _, data := range Query[gogoanime.Watch]("watch/" + id).Sources {
		source = append(source, models.Source{
			Quality: data.Quality,
			Url:     data.Url,
		})
	}

	return source
}

var Package models.Provider = PackageModel{}
