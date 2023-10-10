package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	list := LinkedList{}
	list.sentinel = &Cell{"SENTINEL", nil}
	return list
}

func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}

func (list *LinkedList) addRange(values []string) {
	var lastCell *Cell = list.sentinel
	for ; lastCell.next != nil; lastCell = lastCell.next {
	}

	for _, v := range values {
		newCell := Cell{v, nil}
		lastCell.addAfter(&newCell)
		lastCell = &newCell
	}
}

func (list *LinkedList) toString(seperator string) string {
	return list.toStringMax(seperator, -1)
}

func (list *LinkedList) toStringMax(seperator string, max int) string {
	sb := strings.Builder{}

	cellCount := 0
	var lastCell *Cell = list.sentinel
	totalSize := 0
	var sepLen = len(seperator)
	for ; lastCell.next != nil; lastCell = lastCell.next {
		cellCount++
		if max > 0 && cellCount > max {
			break
		}
		totalSize += len(lastCell.data) + sepLen
	}

	totalSize = totalSize - sepLen
	if totalSize > 0 {
		sb.Grow(totalSize)
	}

	cellCount = 0
	for currCell := list.sentinel.next; currCell != nil; currCell = currCell.next {
		if max > 0 && cellCount >= max {
			break
		}
		sep := seperator
		if currCell.next == nil {
			sep = ""
		}
		sb.WriteString(fmt.Sprintf("%s%s", currCell.data, sep))
		cellCount++
	}

	return sb.String()
}

func (list *LinkedList) hasLoop() bool {
	// https://stackoverflow.com/a/2936345
	nextNode := func(node *Cell) *Cell {
		if node != nil {
			return node.next
		} else {
			return node
		}
	}

	sp := list.sentinel.next
	fp := list.sentinel.next
	for fp.next != nil {
		for m := 0; m < 2; m++ {
			fp = nextNode(fp)
			if fp == nil {
				return false
			}
			if sp == fp {
				return true
			}
		}
		sp = sp.next
	}
	return false
}

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
