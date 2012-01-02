package cappedqueue

import (
	"fmt"
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "cappedqueue: ", log.LstdFlags)
)

type CappedQeuer interface {
	Enqueue(v interface{})
	Dequeue() (v interface{})
	Capacity() int
}

type CappedQueue struct {
	c chan interface{}
}

func New(capacity int) CappedQeuer {
	if capacity == 0 {
		panic("capacity must be greater than 0")
	}
	return &CappedQueue{
		c: make(chan interface{}, capacity),
	}
}

func (c *CappedQueue) Capacity() int {
	return cap(c.c)
}

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

func (c *CappedQueue) Dequeue() interface{} {
	select {
	case v := <-c.c:
		return v
	default:
		return nil
	}
	return nil
}
func Main() {
	fmt.Println("Hi there, I love cappedqueue")
}

// vim: set ts=4 sts=4 fenc=utf-8 noexpandtab list:
