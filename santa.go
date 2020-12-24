'''
You get a list of employees. Create a mapping where each employee is another employees secret santa. 
The mapping should be random on each run of the algorithm. An employee cannot be its own secret santa 
and should not be able to guess its santa knowing the algorithm (i.e. (“Luke”, “Vader”) should not infer (“Vader”, “Luke”))
ppl = ["Vader", "Luke", "Leia", "Han", "Palpatine", "Chewie", "C3PO"]
# example result
result = [('Leia', 'Chewie'), ('Vader', 'Palpatine'), ('Palpatine', 'Leia'), ...]
'''

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
	var employees = []Employe{}
	for i := 0;i < count;i++ {
		name := fmt.Sprintf("%s", i)
		employees = append(employees, Employee{name:name})
	}
	return employees
}


func selectSecretSanta(employees []Employee)  {
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
	for idx, employee := range(employees) {
		employees[idx].skip = true
		employees[idx].secretSanta = selectSecretSanta(employees)
		employees[idx].skip = false
	}

	fmt.Printf("%v\n", employees)
}


