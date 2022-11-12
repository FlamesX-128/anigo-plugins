package main

import (
	"github.com/FlamesX-128/anigo-plugins/consumet/gogoanime"
	"github.com/FlamesX-128/anigo/models"
)

var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"gogoanime": gogoanime.Package,
	},
}
