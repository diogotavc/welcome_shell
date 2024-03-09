package weather

func Report(city string) string {
	return ""
}

func getGlobalIdLocal(city string) int {
	// grab id from https://api.ipma.pt/open-data/distrits-islands.json
	// from a city string - example: Aveiro -> 1010500
}

func getWeatherForecast(id int, days int) [7]string {
	// parse json from https://api.ipma.pt/open-data/forecast/meteorology/cities/daily/{id}.json
	// return array with relevant information
}

// missing same day reports
