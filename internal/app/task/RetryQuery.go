package task

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"unsafe"
)

const defaultConcurrencyNumber = 5
const defaultMaxRetryTimesPerTime = 3

type retryQuery struct {
	rw                   sync.RWMutex
	arena                taskArena
	signal               chan int
	work                 []worker
	nocopy               nocopy
	maxRetryTimesPerTime int32
	workerNumber         int32
	//let query sleep
	remained bool
}
type Option func(retryQuery *retryQuery)

func WithConcurrencyNumber(number int32) Option {
	return func(retryQuery *retryQuery) {
		retryQuery.workerNumber = number
	}
}
func WithMaxRetryTimesPerTime(number int32) Option {
	return func(retryQuery *retryQuery) {
		retryQuery.maxRetryTimesPerTime = number
	}
}

func NewRetryQuery(options ...Option) *retryQuery {
	res := &retryQuery{
		workerNumber:         defaultConcurrencyNumber,
		maxRetryTimesPerTime: defaultMaxRetryTimesPerTime,
		arena:                make(taskArena, 10),
		rw:                   sync.RWMutex{},
		remained:             false,
	}
	res.check()
	for i := 0; i < len(options); i++ {
		options[i](res)
	}
	res.signal = make(chan int, res.workerNumber)
	res.work = make([]worker, res.workerNumber)
	return res
}

type worker struct {
	query      *retryQuery
	task       *task
	state      int32
	retryTimes int32
}

func (work *worker) work() {
	// request work
}

type taskArena []*task

type task struct {
	exec     func() error
	priority int8
}

func (query *retryQuery) check() {
	if !atomic.CompareAndSwapUintptr((*uintptr)(&query.nocopy), uintptr(0), uintptr(unsafe.Pointer(query))) && unsafe.Pointer(query.nocopy) != unsafe.Pointer(query) {
		panic("task copy")
	}
}

type nocopy uintptr

func (arena taskArena) Less(i, j int) bool {
	return arena[i].priority < arena[j].priority
}
func (arena taskArena) Len() int {
	return len(arena)
}
func (arena taskArena) Swap(i, j int) {
	arena[i], arena[j] = arena[j], arena[i]
}

func (arena *taskArena) Push(value interface{}) {
	*arena = append(*arena, value.(*task))
}

func (arena *taskArena) Pop() interface{} {
	old := *arena
	n := len(old)
	x := old[n-1]
	*arena = old[0 : n-1]
	return x
}
func (query *retryQuery) AddTask(function func() error, priority int8) {
	query.check()
	query.rw.Lock()
	defer query.rw.Unlock()
	task := task{
		exec:     function,
		priority: priority,
	}
	heap.Push(&query.arena, task)
}

func (query *retryQuery) Run() {

}
