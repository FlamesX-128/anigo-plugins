package fembed

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo-plugins/monochinos/models/providers/fembed"
)

func Handler(url string) ([]models.Source, error) {
	id := url[(strings.LastIndex(url, "/") + 1):]

	var req *http.Request
	var err error

	if req, err = http.NewRequest("POST", ("https://www.fembed.com/api/source/" + id), nil); err != nil {

		return nil, err
	}

	req.Header.Set("x-requested-with", "XMLHttpRequest")

	var sources []models.Source
	var data fembed.Response
	var resp *http.Response

	if resp, err = http.DefaultClient.Do(req); err != nil {

		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {

		return nil, err
	}

	for _, v := range data.Data {
		sources = append(sources, models.Source(v))
	}

	return sources, nil
}
