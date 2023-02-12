package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	value int
	next  *node
}

type list struct {
	head *node
}

func main() {
	var list1 *list = &list{}

	for i := 0; list1.len() < 5; i++ {
		list1.unique_append(rand.Intn(6))
	}

	list1.print()
	// list1.find()

}

func (l *list) append(val int) {
	if l.head == nil {
		l.head = &node{value: val, next: nil}
	} else {
		var item *node = l.head
		for {
			if item.next == nil {
				item.next = &node{value: val, next: nil}
				break
			} else {
				item = item.next
			}
		}
	}
}

func (l *list) len() int {
	var len int
	if l.head != nil {
		var item node = *l.head
		for {
			len++
			if item.next == nil {
				break
			} else {
				item = *item.next
			}
		}
	}
	return len
}

func (l *list) unique_append(val int) {
	if !(l.find(val)) {
		if l.head == nil {
			l.head = &node{value: val, next: nil}
		} else {
			var item *node = l.head
			for {
				if item.next == nil {
					item.next = &node{value: val, next: nil}
					break
				} else {
					item = item.next
				}
			}
		}
	}
}

func (l *list) print() {
	if l.head != nil {
		var item node = *l.head
		for {
			fmt.Println(item.value)
			if item.next == nil {
				break
			} else {
				item = *item.next
			}
		}
	}
}

func (l *list) find(val int) bool {
	var found bool
	if l.head != nil {
		var item node = *l.head
		for {
			if item.value == val {
				fmt.Printf("The value %v exists in the list\n", val)
				found = true
			}
			if item.next == nil {
				if !found {
					fmt.Printf("The value doesn't exist\n")
					found = false
				}
				break
			} else {
				item = *item.next
			}
		}
	}
	return found
}
