package queue

import "errors"

const (
	defaultQueueSize=11     // max 10 elements
)

var queueSize int

type Queue struct{

	front    int
	rear     int
	currentCount    int
	elements []interface{}
	
}

func NewQueueBySize(size int) *Queue {

	queueSize = size
	return &Queue{0, size-1, 0, make([]interface{},size)}
}

func NewQueue() *Queue{

	return NewQueueBySize(defaultQueueSize)
}

func ProbeNext(i int) int {

	return (i+1)%queueSize
}

func (queue *Queue)ClearQueue(){

	queue.front = 0
	queue.rear = queueSize - 1
	queue.currentCount = 0
}

func (queue *Queue)IsEmpty() bool {

	if ProbeNext(queue.rear) == queue.front {
		return true
	}
	return false
}

func (queue *Queue)IsFull() bool {

	if ProbeNext(ProbeNext(queue.rear)) == queue.front {
		return true
	}

	return false
}

func (queue *Queue)Offer(e interface{}) error {

	if queue.IsFull() ==  true {
		return errors.New("the queue is full.")
	}


	rear := ProbeNext(queue.rear)
	queue.elements[rear] = e
	queue.currentCount = queueSize + 1

	queue.rear = rear

	return nil
}

func (queue *Queue)Poll()(t interface{}, err error) {

	if queue.IsEmpty() == true {
		return nil, errors.New("the queue is empty.")
	}

	front := queue.front
	t = queue.elements[front]
	queue.currentCount = queue.currentCount - 1
	queue.front = ProbeNext(front)

	return t,nil
}