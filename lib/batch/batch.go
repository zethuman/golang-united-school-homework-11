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
	var mx sync.Mutex
	var i int64
	for i = 0; i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			defer wg.Done()
			defer func() { <-sem }()

			user := getOne(j)

			mx.Lock()
			res = append(res, user)
			mx.Unlock()
		}(i)

	}
	wg.Wait()
	return
}
