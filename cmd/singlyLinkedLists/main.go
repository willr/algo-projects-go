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

func (list *LinkedList) appendEnd(appendValue string) {
	var lastCell *Cell = list.sentinel
	for ; lastCell.next != nil; lastCell = lastCell.next {
	}
	newCell := Cell{appendValue, nil}
	lastCell.addAfter(&newCell)
}

func (me *Cell) deleteAfter() {
	if me.next == nil {
		panic("can't delete cell after current. No next cell!")
	}
	me.next = me.next.next
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

func (list *LinkedList) addList(values LinkedList) {
	var lastCell *Cell = list.sentinel
	for ; lastCell.next != nil; lastCell = lastCell.next {
	}

	for valueCell := values.sentinel; lastCell.next != nil; valueCell = valueCell.next {
		lastCell.addAfter(valueCell)
		lastCell = valueCell
	}
}

func (list *LinkedList) toSlice() []string {
	cellCount := 0
	for lastCell := list.sentinel; lastCell.next != nil; lastCell = lastCell.next {
		cellCount++
	}

	valueSlice := make([]string, cellCount)
	currCell := 0
	for lastCell := list.sentinel.next; lastCell != nil; lastCell = lastCell.next {
		valueSlice[currCell] = lastCell.data
		currCell++
	}
	return valueSlice
}

func (list *LinkedList) clone() *LinkedList {
	newList := makeLinkedList()
	newList.sentinel.next = list.sentinel.next
	return &newList
}

func (list *LinkedList) clear() {
	list.sentinel.next = nil
}

func (list *LinkedList) toString(seperator string) string {
	sb := strings.Builder{}

	cellCount := 0
	var lastCell *Cell = list.sentinel
	totalSize := 0
	var sepLen = len(seperator)
	for ; lastCell.next != nil; lastCell = lastCell.next {
		cellCount++
		totalSize += len(lastCell.data) + sepLen
	}

	totalSize = totalSize - sepLen
	if totalSize > 0 {
		sb.Grow(totalSize)
	}

	for currCell := list.sentinel.next; currCell != nil; currCell = currCell.next {
		sep := seperator
		if currCell.next == nil {
			sep = ""
		}
		sb.WriteString(fmt.Sprintf("%s%s", currCell.data, sep))
	}

	return sb.String()
}

func (list *LinkedList) length() int {
	cellCount := 0
	for lastCell := list.sentinel; lastCell.next != nil; lastCell = lastCell.next {
		cellCount++
	}
	return cellCount
}

func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}

func (list *LinkedList) contains(query string) bool {
	found := false
	for lastCell := list.sentinel; lastCell.next != nil; lastCell = lastCell.next {
		if query == lastCell.data {
			found = true
			break
		}
	}
	return found
}

func (list *LinkedList) findBefore(query string) *Cell {
	var foundCell *Cell = nil
	for lastCell := list.sentinel; lastCell.next != nil; lastCell = lastCell.next {
		if lastCell.next == nil {
			break
		}
		if query == lastCell.next.data {
			foundCell = lastCell
			break
		}
	}
	return foundCell
}

func (list *LinkedList) removeByValue(query string) bool {
	beforeCell := list.findBefore(query)
	if beforeCell == nil {
		return false
	}
	afterCell := beforeCell.next.next
	beforeCell.next = afterCell
	return true
}

func (list *LinkedList) removeByPosition(position int) bool {

	if list.isEmpty() {
		panic(fmt.Sprintf("can't delete cell at position: %d, when LinkedList is empty", position))
	}

	if position == 0 {
		if list.sentinel.next != nil {
			list.sentinel.next = list.sentinel.next.next
		} else {
			list.sentinel.next = nil
		}
		return true
	}

	cellCount := 0
	for lastCell := list.sentinel.next; lastCell != nil; lastCell = lastCell.next {
		cellCount++
		if (cellCount) == position {
			lastCell.deleteAfter()
			break
		}
	}
	if position > cellCount {
		panic(fmt.Sprintf("can't delete cell at position: %d, when position is beyond LinkedList length: %d", position, cellCount))
	}

	return true
}

func (list *LinkedList) push(pushedValue string) *Cell {
	newCell := Cell{pushedValue, nil}
	list.sentinel.addAfter(&newCell)
	return &newCell
}

func (list *LinkedList) pop() string {
	if list.isEmpty() {
		panic("can't pop from empty LinkedList")
	}
	poppedValue := list.sentinel.next.data
	list.sentinel.deleteAfter()
	return poppedValue
}

func main() {
	// smallListTest()

	// Make a list from an array of values.
	greekLetters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greekLetters)
	fmt.Println(list.toString(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.toString(" "))
	}
}
