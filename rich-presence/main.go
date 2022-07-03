package main

import (
	"log"
	"reflect"
	"time"

	"github.com/FlamesX-128/anigo-plugins/types"
	"github.com/hugolgst/rich-go/client"
)

var Buttons = []*client.Button{
	{
		Url:   "https://github.com/FlamesX-128/anigo",
		Label: "GitHub",
	},
}

func checkVariables(a *types.Anigo) bool {
	refl := reflect.ValueOf(a)

	for _, e := range []string{"Episode", "Image", "Name", "Url"} {
		if refl.FieldByName(e).String() == "" {

			return false
		}

	}

	return true
}

func Handler(a *types.Anigo) {
	if err := client.Login("940982082891579482"); err != nil {
		log.Println(err)

		return
	}

	var now = time.Now()

	for i := 0; i < 4; {
		time.Sleep(time.Second * 1)

		if !checkVariables(a) {
			now = time.Now()

		}

		err := client.SetActivity(client.Activity{
			State:      "Watching episode " + string(rune(a.Episode)),
			Details:    a.Name,
			LargeImage: a.Image,
			LargeText:  a.Name,
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: Buttons,
		})

		if i == 3 {
			log.Println(err)

			break
		}

		if err != nil {
			i++

		}

	}

}

var Plugin = types.Plugin[types.Process]{
	Author: "FlamesX-128",
	Name:   "rich-presence",
	This: types.Process{
		Handler:    Handler,
		Persistent: false,
	},
	Url:     "github.com/FlamesX-128/anigo-plugins/rich-presence",
	Version: "0.1.0",
}
