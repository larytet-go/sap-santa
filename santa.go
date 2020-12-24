package main

import (
	"fmt"
	"math/rand"
)

type Employee struct {
	name string
	skip bool
	secretSanta int
}

func getEmployes(count int) []Employee {
	var employees = []Employee{}
	for i := 0;i < count;i++ {
		name := fmt.Sprintf("%s", i)
		employees = append(employees, Employee{name:name})
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
	}
	return foundIdx
}
func main() {
	rand.Seed(0) // cutting corners, use tick instead
	employees := getEmployes(7)
	for _, employee := range(employees) {
		employee.skip = true
		employee.secretSanta = selectSecretSanta(employees)
		employee.skip = false
	}

	fmt.Printf("%v\n", employees)
}


