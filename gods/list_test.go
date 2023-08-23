package gods

import (
	"testing"
)

func TestPrependVal(t *testing.T) {
	head := newEntry(10)

	testHeadPrependVal(t, head, 20)
	testHeadPrependVal(t, head, 30)
	testHeadPrependVal(t, head, 40)
}

func testHeadPrependVal(t *testing.T, head *Entry, val any) {
	created := head.PrependVal(val)

	if created == nil {
		t.Fatal("PrependVal() expected to return non-nil pointer")
	}

	if created.Val() != val {
		t.Errorf("PrependVal(%v) = &Entry{Value: %v}, want %[1]v", val, created.Val())
	}

	if created.Prev() != nil {
		t.Error("PrependVal() expected to return first element")
	}
}

func TestAppendVal(t *testing.T) {
	head := newEntry(10)

	testHeadAppendVal(t, head, 20)
	testHeadAppendVal(t, head, 30)
	testHeadAppendVal(t, head, 40)
}

func testHeadAppendVal(t *testing.T, head *Entry, val any) {
	created := head.AppendVal(val)

	if created == nil {
		t.Fatal("AppendVal() expected to return non-nil pointer")
	}

	if created.value != val {
		t.Errorf("AppendVal(%v) = &Entry{Value: %v}, want %[1]v", val, created.Val())
	}

	if created.Next() != nil {
		t.Error("AppendVal() expected to return last element")
	}
}

func TestVal(t *testing.T) {
	head := newEntry(10)

	if head.Val() != 10 {
		t.Errorf("Val() = %v, want %v", head.Val(), 10)
	}
}

func TestPrev(t *testing.T) {
	head := newEntry(10)
	second := head.AppendVal(20)
	tail := second.AppendVal(30)

	if head.Prev() != nil {
		t.Error("expected `head` to be the first entry")
	}

	if second.Prev() != head {
		t.Error("Prev() expected to return `head`")
	}

	if tail.Prev() != second {
		t.Error("`tail`.Prev() expected to return `second`")
	}
}

func TestNext(t *testing.T) {
	head := newEntry(10)
	second := head.AppendVal(20)
	tail := second.AppendVal(30)

	if head.Next() != second {
		t.Error("Next() expected to return second entry")
	}

	if head.Next().Next() != tail {
		t.Error("Next().Next() expected to return third entry")
	}

	if tail.Next() != nil {
		t.Error("expected `tail` to be the last entry")
	}
}

func TestRemove(t *testing.T) {
	// first case: single entry
	head := newEntry(10)

	head.Remove()

	if head.Prev() != nil || head.Next() != nil {
		t.Error("expect removing single entry to do nothing")
	}

	// second case: head
	head = newEntry(10)
	tail := head.AppendVal(20)

	head.Remove()

	if head.Prev() != nil || head.Next() != nil {
		t.Error("expect removing `head` to unlink previous and next entries")
	}

	if tail.Prev() != nil || tail.Next() != nil {
		t.Error("expect removing `head` to let `tail` as a single entry")
	}

	// third case: tail
	head = newEntry(10)
	tail = head.AppendVal(20)

	tail.Remove()

	if tail.Prev() != nil || tail.Next() != nil {
		t.Error("expect removing `tail` to unlink previous and next entries")
	}

	if head.Prev() != nil || head.Next() != nil {
		t.Error("expect removing `tail` to let `head` as a single entry")
	}

	// fourth case: middle entry
	head = newEntry(10)
	middle := head.AppendVal(20)
	tail = middle.AppendVal(30)

	middle.Remove()

	if head.Next() != tail {
		t.Error("expect `head` to point to `tail` forward")
	}

	if tail.Prev() != head {
		t.Error("expect `tail` to point to `head` backward")
	}

	if middle.Prev() == head || middle.Next() == tail {
		t.Error("expect removing `middle` to unlink previous and next entries")
	}
}

func ExampleEntry() {
	head := newEntry(10)
	head.AppendVal(20).AppendVal(30)

	printList(head)

	// Output:
	// 10
	// 20
	// 30
}
