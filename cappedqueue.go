// cappedqueue is a simple package for a queue that never blocks.
//
// Non-Blocking in that case means that enqueued items will be lost once the
// queue is full and more items are submitted.
package cappedqueue

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "cappedqueue: ", log.LstdFlags)
)

// CappedQeuer is the interface that all CappedQueues must implement.
//
// All operations must always succeed and never block.
type CappedQeuer interface {
	Enqueue(v interface{})
	Dequeue() (v interface{})
	Capacity() int
}

type CappedQueue struct {
	c chan interface{}
}

// New creates a new CappedQueue that you can work with.
//
// capacity must be greater than zero or New will panic.
func New(capacity int) CappedQeuer {
	if capacity == 0 {
		panic("capacity must be greater than 0")
	}
	return &CappedQueue{
		c: make(chan interface{}, capacity),
	}
}

// Capacity returns the maximum number of items the CappedQueue can hold
func (c *CappedQueue) Capacity() int {
	return cap(c.c)
}

// Enqueue puts a new item at the end of the queue.
// This operation always succeeds.
//
// If it is not possible to possible to put an item at the end of the queue
// because the Capacity would be exceeded on item will be removed from the
// beginning of the queue and is effectively lost.
func (c *CappedQueue) Enqueue(v interface{}) {
	for {
		select {
		case c.c <- v:
			// If this fires we successfully put v into our queue
			// logger.Printf("Enqueing %#v", v)
			break
		case <-c.c:
			// Our queue was full so we remove an item
			// logger.Printf("Queue Full, dequeued: %#v", v)
			c.c <- v
			return
		}
	}
	return
}

// Dequeue returns the item at the front of the queue or nil if the queue is
// empty.
//
// The Dequeue operation always succeeds but may give you nil if you try to
// deqeue from an empty CappedQueue.
func (c *CappedQueue) Dequeue() interface{} {
	select {
	case v := <-c.c:
		return v
	default:
		return nil
	}
	return nil
}
}

// vim: set ts=4 sts=4 fenc=utf-8 noexpandtab list:
