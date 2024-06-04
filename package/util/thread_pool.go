package util

type ThreadPool struct {
	queue chan func()
}

func (p *ThreadPool) worker() {
	for exec := range p.queue {
		exec()
	}
}

func (p *ThreadPool) Submit(exec func()) {
	p.queue <- exec
}

func NewThreadPool(threadNum int) *ThreadPool {
	p := &ThreadPool{
		queue: make(chan func()),
	}
	for i := 0; i < threadNum; i++ {
		GoSafe(p.worker)
	}
	return p
}
