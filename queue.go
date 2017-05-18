package queue

import (
	"sync"
)

type Queue struct{
	sync.RWMutex
	Events []interface{}
}

// конструктор
func New() (*Queue) {
	return &Queue{}
}

// определение длины очереди
func (q *Queue) Len() int {
	l := len(q.Events)
	return l
}

// вставка новых строк в очередь
func (q *Queue) Push(e interface{}) {
	q.Lock()
	defer q.Unlock()
	q.Events = append(q.Events, e)
}

// вставка новых строк в очередь
func (q *Queue) RollBack(items []interface{}) {
	q.Lock()
	defer q.Unlock()
	q.Events = append(q.Events, items...)
}

// выбираем и обрезаем до заданого размера
func (q *Queue) Pop(count int) ([]interface{}) {
	// лочим очередь на запись чтобы мы могли контролировать удаление элементов
	q.Lock()
	defer q.Unlock()

	var partition []interface{}

	// если кол-во событий достигло нужного числа
	if len(q.Events) >= count {
		// забираем нужное кол-во
		partition = q.Events[:count]
		// обрезаем очередь
		q.Events = q.Events[count:]
	}

	return partition
}