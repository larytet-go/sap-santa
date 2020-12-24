package main

import (
	"fmt"
	"time"
	"http"
	"encoding/json"
)

type Location struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
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
	errJson := json.Unmarshall(body, &locations)
	// I feel lucky!
	locaton := locations[0].Woeid
	return locaton, errJson
}


func selectSecretSanta(employees []Employee) int {
	foundIdx := 0
	for {
		idx := rand.Intn(len(employees))
		employee := employees[idx]
		if employee.skip {
			continue
		}
		foundIdx = idx
		break
	}
	return foundIdx
}

func selectSecretSantas() {
	for idx, _ := range(employees) {
		employee := &employees[idx]
		employee.skip = true
		secretSanta := selectSecretSanta(employees)
		employee.secretSanta = secretSanta
		employee.skip = false
	}

	fmt.Printf("[")
	for idx, employee := range(employees) {
		secretSanta := employee.secretSanta
		secretSantaName := employees[secretSanta].name
		fmt.Printf("('%s', '%s')", employee.name, secretSantaName)
		if idx < len(employees) - 1 {
			fmt.Printf(", ")
		} 
	}
	fmt.Printf("]")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	employees := getEmployees(7)
}


