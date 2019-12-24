package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

const (
	openweather = "http://api.openweathermap.org/data/2.5/weather?q="
	appid       = "appid=d3233ef0b9b6d9bb52b006f9f562e162"
)

type WeatherController struct {
	beego.Controller
}

type Coord struct {
	Lat float32
	Lon float32
}
type Sys struct {
	Type    int
	Country string
	Sunrise int64
	Sunset  int64
}
type Main struct {
	Temp float32
}
type Wind struct {
	Speed float32
	Deg   int
}
type Response struct {
	Coord Coord
	Name  string
	Sys   Sys
	Main  Main
	Wind  Wind
}

// @router /getCurrentWeather [get]
func (c *WeatherController) Get() {
	params := c.Ctx.Request.URL.Query()
	city := params["city"]
	country := params["country"]
	//c.Ctx.Input.Bind(&city, "city")
	req := httplib.Get(openweather + city[0] + "," + country[0] + "&" + appid)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	var r Response
	json.Unmarshal([]byte(str), &r)
	c.Data["json"] = r
	c.ServeJSON()
}
