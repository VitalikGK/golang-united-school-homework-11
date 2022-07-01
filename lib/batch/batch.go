package batch

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	//	m       sync.Mutex
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	g := new(errgroup.Group)
	g.SetLimit(int(pool))
	var m sync.Mutex
	var i int64
	//	var wg sync.WaitGroup

	for i = 0; i < n; i++ {
		index := i

		g.Go(func() error {
			user := getOne(index)

			m.Lock()
			defer m.Unlock()
			fmt.Println("Go index = ", index)
			res = append(res, user)

			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("All Done!", res)
	}

	return res
}
