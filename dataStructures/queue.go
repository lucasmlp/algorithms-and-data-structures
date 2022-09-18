package dataStructures

type queue []string

type Queue interface {
	Enqueue(element string)
	Dequeue() (string, error)
}

func NewQueue() queue {
	return queue{}
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(element string) {
	*q = append(*q, element)
}

func (q *queue) Dequeue() (string, error) {
	if q.isEmpty() {
		return "", ErrQueueEmpty
	}

	element := (*q)[0]
	*q = (*q)[1:]
	return element, nil
}
