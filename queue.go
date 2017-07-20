package queue

import (
	"sync"
)

type Queue struct {
	sync.Mutex
	events []interface{}
}

// конструктор
func New() (*Queue) {
	return &Queue{
		events: make([]interface{}, 0, 10000),
	}
}

// определение длины очереди
func (q *Queue) Close() {
	q.events = make([]interface{}, 0, 10000)
}

// определение длины очереди
func (q *Queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return len(q.events)
}

// вставка новых строк в очередь
func (q *Queue) Push(e interface{}) (length int) {
	q.Lock()
	defer q.Unlock()

	q.events = append(q.events, e)
	length = len(q.events)
	return
}

// выбираем и обрезаем до заданого размера
func (q *Queue) Pop(count int) (partition []interface{}) {
	q.Lock()
	defer q.Unlock()

	if len(q.events) >= count {
		// забираем нужное кол-во
		partition = q.events[:count]
		// обрезаем очередь
		q.events = q.events[count:]
	}
	return
}