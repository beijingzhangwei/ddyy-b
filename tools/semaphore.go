package tools

type slot struct{}
type slots chan slot
type semaphore interface {
	acquire()
	release()
	wait()
	close()
}
type execution struct {
	slots slots
}

func NewExecutionSlots(capacity int) execution {
	slots := make(chan slot, capacity) // capacity = worker count
	return execution{slots: slots}
}
func (e execution) acquire() { // 1 任务执行 获取1信号量
	e.slots <- slot{}
}
func (e execution) release() { // 2 执行结束 释放1信号量
	<-e.slots
}
func (e execution) wait() { // 3 获取所有信号量（如果全部获取成功说明，任务都已经执行完，并把自己的信号量释放掉）
	for i := 0; i < cap(e.slots); i++ {
		e.slots <- slot{} // empty struct type
	}
}

// close the slots channel when done with it
func (e execution) close() {
	close(e.slots)
}
