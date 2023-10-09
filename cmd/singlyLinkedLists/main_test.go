package main

import (
	"fmt"
	"testing"

	"linked-data-structures.bblox.io/internal/assert"
)

func TestMakeLinkedList(t *testing.T) {

	var ll = makeLinkedList()

	const sent_data = "SENTINEL"

	if !(ll.sentinel.data == sent_data) {
		t.Errorf("got %q ; want %q", ll.sentinel.data, sent_data)
	}
}

func TestAddAfter(t *testing.T) {
	var ll = makeLinkedList()
	var newCellData = "A"
	var anotherCellData = "B"
	var newCell = Cell{newCellData, nil}
	var anotherCell = Cell{anotherCellData, nil}

	tests := []struct {
		name        string
		src         *Cell
		addCell     *Cell
		addCellData string
	}{
		{
			name:        "First Cell",
			src:         ll.sentinel,
			addCell:     &newCell,
			addCellData: newCellData,
		},
		{
			name:        "Second Cell",
			src:         &newCell,
			addCell:     &anotherCell,
			addCellData: anotherCellData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.src.addAfter(tt.addCell)

			assert.Equal(t, tt.src.next, tt.addCell)
			assert.Equal(t, tt.src.next.data, tt.addCellData)

		})
	}
}

func TestAppendEnd(t *testing.T) {
	var ll = makeLinkedList()
	var newCellData = "A"
	var anotherCellData = "B"
	var endCellData = "C"
	var newCell = Cell{newCellData, nil}
	var anotherCell = Cell{anotherCellData, nil}

	ll.sentinel.addAfter(&newCell)
	newCell.addAfter(&anotherCell)

	ll.appendEnd(endCellData)

	var lastCell = ll.sentinel.next.next.next
	assert.Equal(t, lastCell.data, endCellData)
	assert.Equal(t, lastCell.next, nil)

}

func TestAddRange(t *testing.T) {
	var newCellData = "A"
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.appendEnd(newCellData)
	ll.addRange(myRange)

	var startCell = ll.sentinel.next

	var count = 0
	for _, v := range myRange {
		count++
		assert.Equal(t, fmt.Sprint(count), v)
		assert.Equal(t, startCell.next.data, v)
		startCell = startCell.next
	}
}

func TestDeleteAfter(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	var llAgain = makeLinkedList()
	llAgain.addRange(myRange)

	var parentCell = ll.sentinel.next.next.next
	parentCell.deleteAfter()
	assert.Equal(t, parentCell.next, nil)

	var p2Cell = ll.sentinel
	p2Cell.deleteAfter()
	assert.Equal(t, p2Cell.next.data, "2")
}

func TestAddList(t *testing.T) {
	var newCellData = "A"
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.appendEnd(newCellData)
	ll.addRange(myRange)

	var basell = makeLinkedList()

	basell.appendEnd(newCellData)
	basell.addList(ll)

	ll.addRange(myRange)

	var count = 0
	for startCell := basell.sentinel.next.next; startCell != nil; startCell = startCell.next {
		var v = startCell.data
		count++
		assert.Equal(t, fmt.Sprint(count), v)
		assert.Equal(t, startCell.next.data, v)
		startCell = startCell.next
	}
}

func TestToSlice(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	var mySlice = ll.toSlice()
	var count = 0
	for i, v := range mySlice {
		count++
		assert.Equal(t, fmt.Sprint(count), v)
		assert.Equal(t, myRange[i], v)
	}
}

func TestClone(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	var newll = ll.clone()
	var mySlice = newll.toSlice()
	var count = 0
	for i, v := range mySlice {
		count++
		assert.Equal(t, fmt.Sprint(count), v)
		assert.Equal(t, myRange[i], v)
	}

	assert.NotEqual(t, &ll, newll)
}

func TestClear(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	ll.clear()

	assert.Equal(t, ll.sentinel.next, nil)
	assert.Equal(t, ll.length(), 0)
}

func TestToString(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	assert.Equal(t, ll.toString(","), "1,2,3,4")
}

func TestLength(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	assert.Equal(t, ll.length(), 4)
}

func TestIsEmpty(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	ll.clear()

	assert.Equal(t, ll.length(), 0)
}

func TestContains(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	assert.Equal(t, ll.contains("3"), true)
	assert.NotEqual(t, ll.contains("SENTINEL"), false)
}

func TestFindBefore(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	var found *Cell = ll.findBefore("3")
	assert.Equal(t, found, ll.sentinel.next.next)
}

func TestRemoveByValue(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)
	ll.removeByValue("1")

	var llAgain = makeLinkedList()
	llAgain.addRange(myRange)
	llAgain.removeByValue("3")

	var llAgain2 = makeLinkedList()
	llAgain2.addRange(myRange)
	llAgain2.removeByValue("4")

	assert.Equal(t, ll.sentinel.next.data, "2")
	assert.Equal(t, llAgain.sentinel.next.next.next.data, "4")
	assert.Equal(t, llAgain2.sentinel.next.next.next.data, "3")
}

func TestRemoveByPosition(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)
	ll.removeByPosition(0)

	var llAgain = makeLinkedList()
	llAgain.addRange(myRange)
	llAgain.removeByPosition(2)

	assert.Equal(t, ll.sentinel.next.data, "2")
	assert.Equal(t, llAgain.sentinel.next.next.next.data, "4")
}

func TestPush(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	ll.push("A")
	beforeCell := ll.findBefore("A")
	assert.Equal(t, beforeCell.data, "SENTINEL")

	var llAgain = makeLinkedList()
	llAgain.push("A")
	beforeCell = llAgain.findBefore("A")
	assert.Equal(t, beforeCell.data, "SENTINEL")
}

func TestPop(t *testing.T) {
	var myRange = []string{"1", "2", "3", "4"}

	var ll = makeLinkedList()
	ll.addRange(myRange)

	headCell := ll.pop()
	assert.Equal(t, headCell, "1")
	assert.Equal(t, ll.sentinel.next.data, "2")
	assert.Equal(t, ll.length(), 3)
}
