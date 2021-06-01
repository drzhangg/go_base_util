package demo3

import (
	"fmt"
	"time"
)

type worker struct {
	id  int
	err error
}

func (wk *worker) work(workerChan chan<- *worker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("Panic happend with [%v]", r)
			}
		} else {
			wk.err = err
		}

		//通知主goroutine，当前子goroutine已经死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start worker ... ID=",wk.id)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
	}

	panic("worker panic...")
	return err
}
