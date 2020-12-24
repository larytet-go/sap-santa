package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
	"strings"
)

type Location struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}


type Weather struct {
	ID                   int64     `json:"id"`
	WeatherStateName     string    `json:"weather_state_name"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	Created              time.Time `json:"created"`
	ApplicableDate       string    `json:"applicable_date"`
	MinTemp              float64   `json:"min_temp"`
	MaxTemp              float64   `json:"max_temp"`
	TheTemp              float64   `json:"the_temp"`
	WindSpeed            float64   `json:"wind_speed"`
	WindDirection        float64   `json:"wind_direction"`
	AirPressure          float64   `json:"air_pressure"`
	Humidity             int       `json:"humidity"`
	Visibility           float64   `json:"visibility"`
	Predictability       int       `json:"predictability"`
}

// Do something like this
// curl https://www.metaweather.com/api/location/search/?query=berlin
// cache goes here if needed
func getLocation(name string) (int, error) {
	query := fmt.Sprintf("https://www.metaweather.com/api/location/search/?query=%s", name)
	resp, errGet := http.Get(query)
	if errGet != nil {
		return 0, errGet
	}
	defer resp.Body.Close()
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return 0, errBody
	}
	locations := []Location{}
	errJson := json.Unmarshal(body, &locations)
	// I feel lucky!
	if len(locations) < 1 {
		return 0, fmt.Errorf("no location %s", name)
	}
	locaton := locations[0].Woeid
	fmt.Printf("%v\n", locations[0])
	return locaton, errJson
}

// Something like curl https://www.metaweather.com//api/location/638242/2013/4/27/
func getCleanDay(days int, location int) (time.Time, error) {
	dateToCheck := time.Now()
	year, month, day := dateToCheck.Date()
	query := fmt.Sprintf("https://www.metaweather.com//api/location/%d/%d/%d/%d/", location, year, month, day)
	resp, errGet := http.Get(query)
	if errGet != nil {
		return dateToCheck, errGet
	}
	defer resp.Body.Close()
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return dateToCheck, errBody
	}
	weather := []Weather{}
	errJson := json.Unmarshal(body, &weather)
	if errJson != nil {
		return dateToCheck, fmt.Errorf("%v %s", errJson, string(body))
	}
	if days > 6 {
		days = 6
	}

	for i := 0;i < days;i++ {
		dayWeather := weather[i]
		fmt.Printf("weather %v\n", dayWeather.WeatherStateName)
	}

	for i := 0;i < days;i++ {
		dayWeather := weather[i]
		if strings.Contains(dayWeather.WeatherStateName, "Clear") {
			return dateToCheck, nil
		} else {
			fmt.Printf("weather %v\n", dayWeather.WeatherStateName)
		}
		dateToCheck = dateToCheck.Add(24*time.Hour)
	}
	return dateToCheck, fmt.Errorf("No clean day in the next %d days", days)
}

func main() {
	location, err := getLocation("Perth")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	cleanDay, errGetDay := getCleanDay(5, location)
	if errGetDay != nil {
		fmt.Printf("%v", errGetDay)
		os.Exit(-2)
	}
	fmt.Printf("%v", cleanDay)
}


