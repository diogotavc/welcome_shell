package weather

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Location struct {
	IdRegiao      int    `json:"idRegiao"`
	IdAreaAviso   string `json:"idAreaAviso"`
	IdConcelho    int    `json:"idConcelho"`
	GlobalIdLocal int    `json:"globalIdLocal"`
	Latitude      string `json:"latitude"`
	IdDistrito    int    `json:"idDistrito"`
	Local         string `json:"local"`
	Longitude     string `json:"longitude"`
}

type Forecast struct {
	PrecipitaProb  string `json:"precipitaProb"`
	TMin           string `json:"tMin"`
	TMax           string `json:"tMax"`
	PredWindDir    string `json:"predWindDir"`
	IdWeatherType  int    `json:"idWeatherType"`
	ClassWindSpeed int    `json:"classWindSpeed"`
	Longitude      string `json:"longitude"`
	ForecastDate   string `json:"forecastDate"`
	ClassPrecInt   int    `json:"classPrecInt"`
	Latitude       string `json:"latitude"`
}

type IPMA struct {
	Owner   string     `json:"owner"`
	Country string     `json:"country"`
	Data    []Location `json:"data"`
}

func Report(city string) string {
	forecast := getWeatherForecast(getGlobalIdLocal(city), 1)
	forecastJSON, err := json.Marshal(forecast)
	checkError(err)
	return string(forecastJSON)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getGlobalIdLocal(city string) int {
	resp, err := http.Get("https://api.ipma.pt/open-data/distrits-islands.json")
	checkError(err)
	defer resp.Body.Close()

	var locations IPMA
	err = json.NewDecoder(resp.Body).Decode(&locations)
	checkError(err)

	for _, location := range locations.Data {
		if location.Local == city {
			return location.GlobalIdLocal
		}
	}
	return 0
}

func getWeatherForecast(id int, days int) IPMA {

	resp, err := http.Get("https://api.ipma.pt/open-data/forecast/meteorology/cities/daily/" + strconv.Itoa(id) + ".json")
	checkError(err)
	defer resp.Body.Close()

	var forecasts IPMA
	err = json.NewDecoder(resp.Body).Decode(&forecasts)
	checkError(err)

	return forecasts
}

// missing same day reports
