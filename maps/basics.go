package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	ages["richard"] = 28
	ages["daniel"] = 31
	ages["michael"] = 45

	for k, v := range ages {
		fmt.Printf(k)
		fmt.Println(v)

	}

	males := make(map[string]int)

	males["Arnold"] = 19
	males["Tim"] = 29
	males["Joseph"] = 41
	males["Richard"] = 28

	// check if males and and ages have same values

	for k, v := range males {
		if _, ok := ages[k]; ok {
			if v == ages[k] {
				fmt.Println(k, " exists in both maps and ages is ", ages[k])
			} else {
				fmt.Println("age of ", k, "is different in ages and males")
			}

		} else {
			fmt.Println(k, " doesn't exist in ages")
		}
	}
}
