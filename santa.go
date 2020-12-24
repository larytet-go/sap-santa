

package main

import (
	"fmt"
)


func getEmployes() []string {
	var employees = []string{}
	for i := 0;i < 7;;i++{
		name := fmt.Sprintf("%s", i)
		employees = append(employees, name)
		}
	return employees
}
func main() {
}


