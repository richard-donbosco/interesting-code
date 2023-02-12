package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	prev  *node
	next  *node
	value int
}

type list struct {
	head *node
	tail *node
}

func main() {
	var list1 *list = &list{}

	//append to list
	if val, ok := list1.append(1); ok {
		fmt.Printf("The value %v was appended to the list \n", val)
	}

	list1.print()
	fmt.Println("---------")

	// clear list
	list1.head = nil
	list1.tail = nil

	// remove val from list
	list1.append(1)
	list1.append(2)
	list1.append(3)
	list1.append(4)
	list1.append(5)

	list1.print()
	fmt.Println("---------")

	list1.remove(1)
	list1.print()
	fmt.Println("---------")
	list1.remove(5)
	list1.print()
	fmt.Println("---------")
	list1.remove(3)
	list1.print()
	fmt.Println("---------")

	//get length
	fmt.Println(list1.len())
	fmt.Println("---------")

	//prepend
	list1.prepend(1000)
	list1.prepend(1000)
	list1.print()
	fmt.Println("---------")

	list1.reverse_print()
	fmt.Println("---------")

	fmt.Println(list1.smallest())
	fmt.Println("---------")

	//clear list
	list1.head = nil
	list1.tail = nil

	//sort
	for list1.len() < 100 {
		list1.unique_append(rand.Intn(100))
	}
	list1.print()
	fmt.Println("---------")

	list1.sort()
	list1.print()
	fmt.Println("---------")

	//clear list
	list1.head = nil
	list1.tail = nil

	//sort
	for list1.len() < 100 {
		list1.unique_prepend(rand.Intn(100))
	}
	list1.print()
	fmt.Println("---------")

	list1.sort()
	list1.print()
	fmt.Println("---------")

	//clear list
	list1.head = nil
	list1.tail = nil

	//sort
	for list1.len() < 100 {
		list1.insert(rand.Intn(100))
	}
	list1.print()
	fmt.Println("---------")

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

func (l *list) exists(val int) (int, bool) {
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
	return val, found
}

func (l *list) append(val int) (int, bool) {
	if l.head == nil {
		l.head = &node{prev: nil, next: nil, value: val}
		l.tail = l.head
		return val, true
	}

	if l.tail == l.head {
		l.tail = &node{prev: l.head, next: nil, value: val}
		l.head.next = l.tail
		return val, true
	} else {
		l.tail.next = &node{prev: l.tail, next: nil, value: val}
		l.tail = l.tail.next
		return val, true
	}
}

func (l *list) unique_append(val int) (int, bool) {
	if _, ok := l.exists(val); !ok {
		val, ok := l.append(val)
		return val, ok
	}
	return 0, false
}

func (l *list) prepend(val int) (int, bool) {
	if l.head == nil {
		l.head = &node{next: nil, prev: nil, value: val}
		l.tail = l.head
		return val, true
	} else {
		var tmp *node = l.head
		l.head = &node{next: tmp, prev: nil, value: val}
		tmp.prev = l.head
		return val, true
	}
}

func (l *list) unique_prepend(val int) (int, bool) {
	if _, ok := l.exists(val); !ok {
		val, ok := l.prepend(val)
		return val, ok
	}
	return 0, false
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

func (l *list) reverse_print() {
	if l.tail != nil {
		var item node = *l.tail
		for {
			fmt.Println(item.value)
			if item.prev == nil {
				break
			} else {
				item = *item.prev
			}
		}
	} else {
		fmt.Println("The list is empty")
	}
}

func (l *list) remove(val int) (int, bool) {
	if l.head == nil {
		return 0, true
	}
	if l.head.value == val {
		if l.head.next != nil {
			l.head = l.head.next
		} else {
			l.head = nil
		}
		return val, true
	}
	var current *node = l.head.next
	for {
		if current.value == val && l.tail != current {
			current.prev.next = current.next
			current.next.prev = current.prev
			return val, false
		} else if current.value == val && l.tail == current {
			current.prev.next = nil
			l.tail = current.prev
			return val, false
		}
		if current.next != nil {
			current = current.next
		} else {
			break
		}
	}
	return 0, false

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
		list2.append(val)
		l.remove(val)
	}
	l.head = list2.head
	l.tail = list2.tail
}

func (l *list) insert(val int) (int, bool) {
	if l.len() == 0 {
		l.head = &node{prev: nil, next: nil, value: val}
		l.tail = l.head
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
			current.next = &node{prev: current, next: tmp, value: val}
			tmp.prev = current.next
			return val, true
		} else if current.next != nil {
			current = current.next
		} else if current.next == nil {
			break
		}
	}
	return 0, false
}
