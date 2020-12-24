package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Employee struct {
	name string
	skip bool
	secretSanta int
}

func getEmployes(count int) []Employee {
	ppl := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}
	var employees = []Employee{}
	for i := 0;i < count;i++ {
		name := ppl[i]
		employee := Employee{
			name:name,
		}
		employees = append(employees, employee)
	}
	return employees
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
func main() {
	rand.Seed(time.Now().UnixNano())
	employees := getEmployes(7)
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


