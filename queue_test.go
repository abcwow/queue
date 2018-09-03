package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	
	q := NewQueue()

	if q.IsEmpty() != true {

		t.Error("NewQueue() failed: Not empty, but Empty queue expected.")
	}
}

func TestOffer(t *testing.T) {

	q := NewQueue()

	s := "hello, queue."
	q.Offer(s)
	d, _ := q.Poll()
	if d.(string) != s {

		t.Errorf("Queue.Poll() failed: Got %v, expected %s", d, s)
	}
}

func TestIsFull(t *testing.T) {

	
	q := NewQueue()

	for i:=0; i<10; i++ {

		err := q.Offer(i)
		if err != nil {

			t.Errorf("Queue.Offer() failed: %s", err.Error())
		}
		
	}

	if q.IsFull() != true {

		t.Error("Queue.IsFull() failed: Not full, but Full queue is expected")
	}
	
}

func TestPoll(t *testing.T) {

	q := NewQueue()

	for i:=0; i<10; i++ {
		
		q.Offer(i)
	}

	for i:=0; i<10; i++ {

		d, _ := q.Poll()
		if d != i {
			t.Errorf("Queue.Poll() failed: Got %d, but %d expected", d, i)
		}
	}
}