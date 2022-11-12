package main

import (
	"github.com/FlamesX-128/anigo-plugins/consumet/gogoanime"
	"github.com/FlamesX-128/anigo-plugins/models"
)

var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"[API] [Direct] gogoanime": gogoanime.Package,
	},
}
