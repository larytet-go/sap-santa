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
)


func getEmployes() []string {
	var employees = []string{}
	for (i := 0;i < 7;;i++ {
		name := fmt.Sprintf("%s", i)
		employees = append(employees, name)
		}
	return employees
}


func selectSecrectSanta
func main() {
}


