package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Employee struct {
	name string
	secretSanta string
}

func getEmployees(count int) []Employee {
	ppl := []string{"Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"}
	var employees = make([]Employee, len(ppl))
	for idx, name := range(ppl) {
		employees[idx] = Employee{
			name:name,
		}
	}
	return employees
}

func shuffleEmployees(employees []Employee) {
	for i := 0;i < len(employees);i++ {
		tmp := employees[i]
		j := rand.Intn(len(employees))
		employees[i] = employees[j]
		employees[j] = tmp
	}
}

func selectSecretSantas(employees []Employee) {
	if len (employees) < 2 {
		return
	}
	for idx, _ := range(employees) {
		if idx < len(employees) - 1 {
			employees[idx].secretSanta = employees[idx+1].name
		} else {
			employees[idx].secretSanta = employees[idx-1].name
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	employees := getEmployees(7)
	shuffleEmployees(employees)
	selectSecretSantas(employees)

	fmt.Printf("[")
	for idx, employee := range(employees) {
		secretSanta := employee.secretSanta
		fmt.Printf("('%s', '%s')", employee.name, secretSanta)
		if idx < len(employees) - 1 {
			fmt.Printf(", ")
		} 
	}
	fmt.Printf("]\n")
}


