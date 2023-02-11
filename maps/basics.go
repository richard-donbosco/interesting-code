package main

import (
	"fmt"
	"strings"
)

func main() {
	ages := make(map[string]int)

	add_to_map(ages, "richard", 28)
	add_to_map(ages, "daniel", 31)
	add_to_map(ages, "michael", 45)

	for k, v := range ages {
		fmt.Println(k, " ", v)
	}

	males := make(map[string]int)

	add_to_map(males, "Arnold", 19)
	add_to_map(males, "Tim", 29)
	add_to_map(males, "Joseph", 41)
	add_to_map(males, "Richard", 28)

	// check if males and and ages have same values

	for k, v := range males {
		if _, ok := ages[k]; ok {
			if v == ages[k] {
				fmt.Println(k, " exists in both maps and ages, value is ", ages[k])
			} else {
				fmt.Println("age of ", k, " is different in ages and males")
			}

		} else {
			fmt.Println(k, " doesn't exist in ages")
		}
	}

}

func add_to_map(a map[string]int, k string, v int) {
	k = strings.ToLower(k)
	if _, ok := a[k]; !ok {
		a[k] = v
	} else {
		fmt.Println(k, " already exists in the map")
	}
}
