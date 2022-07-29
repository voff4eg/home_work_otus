package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNoWorker            = errors.New("need at least one worker")
	ErrLimitErrors         = errors.New("m should be more zero")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrLimitErrors
	}

	if n <= 0 {
		return ErrNoWorker
	}

	var errCnt int32
	wg := sync.WaitGroup{}
	wg.Add(n)
	ch := make(chan Task)

	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for task := range ch {
				if e := task(); e != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}(&wg)
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}

		ch <- task
	}

	close(ch)

	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
