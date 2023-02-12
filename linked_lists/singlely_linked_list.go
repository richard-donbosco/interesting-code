package main

import "fmt"

type node struct {
	value int
	next  *node
}

type list struct {
	head *node
}

func main() {
	var list1 *list = &list{}

	for i := 0; i < 10000; i++ {
		list1.append(i)
	}

	list1.print()

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
