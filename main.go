//package notmain

import()

type weatherData struct {
	Location    string `json:"location"`
	Weather     string `json:"weather"`
	Temperature string `json:"temperature"`
	Celsius     string `json:"celsius"`
	Forecast    []int  `json:"forecast"`
	Wind			  windData `json:wind`
}

type windData struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request){

	


}