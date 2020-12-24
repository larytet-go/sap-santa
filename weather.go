package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"strings"
)

type Location struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}

type Weather struct {
	ConsolidatedWeather []struct {
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
	} `json:"consolidated_weather"`
	Time         time.Time `json:"time"`
	SunRise      time.Time `json:"sun_rise"`
	SunSet       time.Time `json:"sun_set"`
	TimezoneName string    `json:"timezone_name"`
	Parent       struct {
		Title        string `json:"title"`
		LocationType string `json:"location_type"`
		Woeid        int    `json:"woeid"`
		LattLong     string `json:"latt_long"`
	} `json:"parent"`
	Sources []struct {
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
		CrawlRate int    `json:"crawl_rate"`
	} `json:"sources"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
	Timezone     string `json:"timezone"`
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
	locaton := locations[0].Woeid
	return locaton, errJson
}

// Something like curl https://www.metaweather.com//api/location/638242/2013/4/27/
func getCleanDay(days int, location int) (time.Time, error) {
	dateToCheck := time.Now()
	year, month, day := dateToCheck.Date()
	query := fmt.Sprintf("https://www.metaweather.com//api/location/%d/%d/%d/%d/", location, year, month, day)
	dateToCheck := dateToCheck.Add(24*time.Hour)
	resp, errGet := http.Get(query)
	if errGet != nil {
		return dateToCheck, errGet
	}
	defer resp.Body.Close()
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return dateToCheck, errBody
	}
	weather := Weather{}
	errJson := json.Unmarshal(body, &weather)
	if errJson != nil {
		return dateToCheck, errJson
	}
	if days > 6 {
		days = 6
	}
	for i := 0;i < days;i++ {
		consolidatedWeather := weather.ConsolidatedWeather[i]
		if strings.Contains(consolidatedWeather.WeatherStateName, "Clear") {
			return dateToCheck, nill
		}
	}
	return dateToCheck, fmt.Errorf("No clean day in the next %d days", days)
}

func main() {
	location, err := getLocation("Berlin")
	if err != nil {
		log.Fatalf("%v", err)
	}
	cleanDay, errGetDay := getCleanDay()
	if errGetDay != nil {
		log.Fatalf("%v", errGetDay)
	}
	log.Infof("%v", cleanDay)
}


