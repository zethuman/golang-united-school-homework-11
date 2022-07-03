package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	mx := &sync.RWMutex{}
	var i int64
	wg.Add(int(n))
	for i = 0; i < n; i++ {
		sem <- struct{}{}
		go func(j int64, mu *sync.RWMutex) {
			defer wg.Done()
			defer func() { <-sem }()
			mu.Lock()
			res = append(res, getOne(j))
			mu.Unlock()
		}(i, mx)

	}
	wg.Wait()
	return
}
