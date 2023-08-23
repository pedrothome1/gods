package gods

import "fmt"

type Entry struct {
	value    any
	previous *Entry
	next     *Entry
}

func (x *Entry) PrependVal(val any) *Entry {
	entry := x

	for entry.previous != nil {
		entry = entry.previous
	}

	entry.previous = &Entry{value: val, next: entry}

	return entry.previous
}

func (x *Entry) AppendVal(val any) *Entry {
	entry := x

	for entry.next != nil {
		entry = entry.next
	}

	entry.next = &Entry{value: val, previous: entry}

	return entry.next
}

func (x *Entry) Remove() {
	// first case: single entry
	if x.previous == nil && x.next == nil {
		return
	}

	// second case: head
	if x.previous == nil {
		x.next.previous = nil
		x.next = nil

		return
	}

	// third case: tail
	if x.next == nil {
		x.previous.next = nil
		x.previous = nil

		return
	}

	// fourth case: middle entry
	x.previous.next = x.next
	x.next.previous = x.previous
	x.next = nil
	x.previous = nil
}

func (x *Entry) Val() any {
	return x.value
}

func (x *Entry) Prev() *Entry {
	return x.previous
}

func (x *Entry) Next() *Entry {
	return x.next
}

func newEntry(value any) *Entry {
	return &Entry{value: value}
}

func printList(head *Entry) {
	for head != nil {
		fmt.Println(head.value)
		head = head.Next()
	}
}
