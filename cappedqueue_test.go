package cappedqueue

import (
	"testing"
)

func TestNew(t *testing.T) {
	New(0)
}

// func TestEmptyDequeue(t *testing.T){
//   cq := New(1)
//   item := cq.Dequeue()
//   t.Logf("ITEM: ", item)
// }

func TestCapacity(t *testing.T) {
	cq := New(1)

	l := cq.Capacity()

	if l != 1 {
		t.Errorf("Got %d, expected 1", l)
	}

}

func TestFull(t *testing.T) {
	cq := New(1)

	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)

	item := cq.Dequeue()

	if item != 3 {
		t.Errorf("Got %d, expected 3", item)
	}
}
