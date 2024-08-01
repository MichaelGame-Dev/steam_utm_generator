package main

import (
	"fmt"
)

// https://store.steampowered.com/app/2693660/Radiant_Bricks?utm_source=website&utm_campaign=summer_sale&utm_medium=web
//    utm_source
//    utm_campaign
//    utm_medium
//    utm_content
//    utm_term

func main() {
	type SteamURL struct {
		baseURL  string
		game     string
		source   string
		campaign string
		medium   string
		content  string
		term     string
	}

	radiantBricks := SteamURL{
		baseURL: "https://store.steampowered.com/app/",
		game:    "2693660/Radiant_Bricks",
		source:  "reddit"}

	utm_url := radiantBricks.baseURL + radiantBricks.game + "?"
	if radiantBricks.source != "" {
		utm_url = utm_url + "utm_source=" + radiantBricks.source
	}
	if radiantBricks.campaign != "" {
		utm_url = utm_url + "utm_source=" + radiantBricks.campaign
	}
	if radiantBricks.source != "" {
		utm_url = utm_url + "utm_source=" + radiantBricks.source
	}
	//var gameURL = url.Parse(radiantBricks.baseURL + radiantBricks.game)
	fmt.Println(utm_url)

}
