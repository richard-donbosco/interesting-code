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

	map_print(ages)
	males := make(map[string]int)

	add_to_map(males, "Arnold", 19)
	add_to_map(males, "Tim", 29)
	add_to_map(males, "Joseph", 41)
	add_to_map(males, "Richard", 28)

	map_print(males)

	// check if males and and ages have same values
	var matching_keys map[string]int
	matching_keys = maps_intersect(males, ages)
	map_print(matching_keys)

	var union_map map[string]int = maps_union(males, ages)
	map_print(union_map)
}

func map_print(a map[string]int) {
	fmt.Println("key\tvalue")
	for k, v := range a {
		fmt.Println(fmt.Sprintf("%s\t%d", k, v))
	}
}

func map_print_matching_keypairs(a map[string]int, keys []string) {
	fmt.Println("key\tvalue")
	for k, v := range a {
		if slice_contains(keys, k) {
			fmt.Println(fmt.Sprintf("%s\t%d", k, v))
		}
	}
}

func slice_contains(a []string, value string) bool {
	for _, v := range a {
		if v == value {
			return true
		}
	}
	return false
}

func add_to_map(a map[string]int, k string, v int) {
	k = strings.ToLower(k)
	if _, ok := a[k]; !ok {
		a[k] = v
	} else {
		fmt.Println(k, " already exists in the map")
	}
}

func delete_from_map(a map[string]int, k string) {
	k = strings.ToLower(k)
	if _, ok := a[k]; !ok {
		delete(a, k)
	} else {
		fmt.Println(k, " doesn't exist in the map")
	}
}

func maps_intersect(a, b map[string]int) map[string]int {
	matching_keys := make(map[string]int)
	for k, v := range a {
		if _, ok := b[k]; ok {
			if v == b[k] {
				matching_keys[k] = v
				fmt.Println(k, " exists in both maps and ages, value is ", b[k])
			} else {
				fmt.Println("age of ", k, " is different in ages and males")
			}

		} else {
			fmt.Println(k, " doesn't exist in ages")
		}
	}
	return matching_keys
}

func maps_union(a, b map[string]int) map[string]int {
	union_map := make(map[string]int)
	for k, v := range a {
		add_to_map(union_map, k, v)
	}
	for k, v := range b {
		add_to_map(union_map, k, v)
	}
	return union_map
}
