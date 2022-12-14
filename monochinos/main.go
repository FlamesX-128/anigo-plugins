package main

import (
	"github.com/FlamesX-128/anigo-plugins/models"
	monoschinos2 "github.com/FlamesX-128/anigo-plugins/monochinos/monochinos2"
)

var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"[Scrape] Monochinos2": monoschinos2.Package,
	},
}
