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
	return &CappedQueue{
		c: make(chan interface{}, capacity),
	}
}

func (c *CappedQueue) Capacity() int {
	return cap(c.c)
}

func (c *CappedQueue) Enqueue(v interface{}) {
	select {
	case c.c <- v:
		logger.Printf("Enqueing %#v", v)
		return
	case v := <-c.c:
		logger.Printf("Queue Full, dequeued: %#v", v)
		c.c <- v
		return
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
