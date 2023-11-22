package queue

type Queue []interface{}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}
