package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3))

	fmt.Println(join_string(",", "richard", "michael", "donbosco", "daniel", "david"))

	var names_list = []string{"richard", "parker", "peter", "logan"}
	fmt.Println(join_string(",", names_list...))

}

func sum(values ...int) int {
	var sum int = 0
	for _, val := range values {
		sum = sum + val
	}
	return sum
}

func join_string(del string, val ...string) string {
	var line string
	for i, val := range val {
		line = line + val
		if i != len(val)-1 {
			line = line + del + " "
		}
	}

	return line

}
