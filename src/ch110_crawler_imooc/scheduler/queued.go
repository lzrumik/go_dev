package scheduler

import "ch110_crawler_imooc/engine"

//调度器结构体
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request  //将管道放入管道
}

//提交request 到 队列
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//管道里面嵌套管道
/**

 */
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

/**

 */
func (s *QueuedScheduler) ConfigureMasterWorkerChannel(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan =make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequet  engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequet = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <-activeRequet:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()

}
