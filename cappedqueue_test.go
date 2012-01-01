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

const (
	maxCapacity = 10000
)

func TestCapacity(t *testing.T) {

	for capacity := 1; capacity < maxCapacity; capacity++ {
		t.Logf("Running with cap of %d", capacity)
		cq := New(capacity)
		got := cq.Capacity()

		if got != capacity {
			t.Errorf("Got %d, expected 1", got)
		}
	}

}

func TestFull(t *testing.T) {
	for capacity := 1; capacity < maxCapacity; capacity++ {
		expected := capacity * 10
		cq := New(capacity)
		t.Logf("Running with cap of %d and expecting %d", capacity,
			expected)
		for i := 1; i < (expected + capacity); i++ {
			cq.Enqueue(i)
		}
		got := cq.Dequeue()
		if got != expected {
			t.Fatalf("Got %d, expected %d", got, expected)
		}
	}
}
