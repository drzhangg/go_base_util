package demo3

import "fmt"

type WorkerManager struct {
	//用来监控worker是否已经死亡的缓冲channel
	workChan chan *worker
	//一共要监控的worker数量
	nWorkers int
}

//实例化对象
func NewWorkerManager(nworkers int) *WorkerManager {
	return &WorkerManager{
		workChan: make(chan *worker, nworkers),
		nWorkers: nworkers,
	}
}

//启动worker池，并为每个worker分配一个id，让每个worker进行工作
func (wm *WorkerManager) StartWorkerPool() {
	for i := 0; i < wm.nWorkers; i++ {
		wk := worker{id: i}
		go wk.work(wm.workChan)
	}

	//启动保活监控
	wm.KeepLiveWorkers()
}

//保活监控workers
func (wm *WorkerManager) KeepLiveWorkers() {
	for wk := range wm.workChan {
		fmt.Printf("Worker %d stopped with err: [%v] \n", wk.id, wk.err)

		wk.err = nil

		// 当前这个wk已经死亡了，需要重新启动他的业务
		go wk.work(wm.workChan)
	}
}
