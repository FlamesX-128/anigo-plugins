package providers

import (
	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo-plugins/monochinos/controllers/providers/fembed"
)

var providersHandler = map[string]func(string) ([]models.Source, error){
	"fembed": fembed.Handler,
}

func ProviderHandler(providers []models.Source) ([]models.Source, bool) {
	for _, value := range providers {
		if handler, ok := providersHandler[value.Quality]; ok {
			if sources, err := handler(value.Url); err == nil {
				return sources, true
			}
		}
	}

	return nil, false
}
