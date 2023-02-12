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

	//append large number of random values
	for i := 0; i < 100; i++ {
		list1.unique_append(rand.Intn(1000))
	}

	list1.print()
	fmt.Println("---------")

	list1.sort()
	list1.print()
	fmt.Println("---------")

	//clear list
	list1.head = nil

	//append large number of random values
	for i := 0; i < 10; i++ {
		list1.insert(rand.Intn(100))
	}
	list1.print()
	fmt.Println("---------")

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
func (l *list) prepend(val int) {
	if l.head == nil {
		l.head = &node{value: val, next: nil}
	} else {
		var tmp *node = l.head
		l.head = &node{value: val, next: tmp}
	}
}
func (l *list) unique_prepend(val int) {
	if !(l.exists(val)) {
		l.prepend(val)
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
	if !(l.exists(val)) {
		l.append(val)
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
	} else {
		fmt.Println("The list is empty")
	}
}

func (l *list) exists(val int) bool {
	var found bool
	if l.head != nil {
		var item node = *l.head
		for {
			if item.value == val {
				// fmt.Printf("The value %v exists in the list\n", val)
				found = true
			}
			if item.next == nil {
				if !found {
					// fmt.Printf("The value doesn't exist\n")
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

func (l *list) smallest() (int, bool) {
	if l.len() == 0 {
		fmt.Println("The list is empty")
		return 0, false
	}
	var smallest int
	if l.head != nil {
		var item node = *l.head
		smallest = item.value
		for {
			if item.next != nil {
				if item.next.value < smallest {
					smallest = item.next.value
				}
			} else {
				break
			}
			item = *item.next
		}
	}
	return smallest, true
}

func (l *list) sort() {
	var list2 *list = &list{}
	var length int = l.len()
	for i := 1; i <= length; i++ {
		val, _ := l.smallest()
		list2.unique_append(val)
		l.remove(val)
	}
	l.head = list2.head
}

func (l *list) remove(val int) (int, bool) {
	if l.head == nil {
		return 0, false
	}
	if l.head.value == val {
		if l.head.next != nil {
			l.head = l.head.next
		} else {
			l.head = nil
		}
		return val, true
	}

	var current *node = l.head
	for {
		if current.next != nil {
			if current.next.value == val {
				if current.next.next != nil {
					current.next = current.next.next
				} else {
					current.next = nil
				}
				return val, true
			} else {
				current = current.next
			}
		} else {
			break
		}
	}

	return 0, false

}

func (l *list) insert(val int) (int, bool) {
	if l.len() == 0 {
		l.head = &node{next: nil, value: val}
		return val, true
	}
	var current *node = l.head
	for {
		if val < current.value {
			l.unique_prepend(val)
			return val, true
		} else if val > current.value && current.next == nil {
			l.unique_append(val)
			return val, true
		} else if val > current.value && val < current.next.value {
			var tmp *node = current.next
			current.next = &node{next: tmp, value: val}
			return val, true
		} else if current.next != nil {
			current = current.next
		} else if current.next == nil {
			break
		}
	}
	return 0, false
}
