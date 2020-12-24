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
		name := fmt.Sprintf("%v", i)
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
	fmt.Printf("foundIdx=%d\n", foundIdx)
	return foundIdx
}
func main() {
	rand.Seed(0) // cutting corners, use tick instead
	employees := getEmployes(7)
	for idx, _ := range(employees) {
		employee := &employees[idx]
		employee.skip = true
		secretSanta := selectSecretSanta(employees)
		employee.secretSanta = secretSanta
		employee.skip = false
	}

	for idx, employee := range(employees) {
		fmt.Printf("%d %s %d\n", idx, employee.name, employee.secretSanta)
	}
}


