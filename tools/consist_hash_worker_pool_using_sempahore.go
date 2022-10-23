package tools

import (
	"fmt"
)

type hardWorker struct {
	selfQueue chan JobInput // 自己的队列
}

func NewHardWorker(maxCap int) *hardWorker {
	selfQueue := make(chan JobInput, maxCap) // 私有队列
	return &hardWorker{selfQueue: selfQueue}
}

func (h *hardWorker) doWorkUsingSelfQueue(jobOutputChan chan<- JobResponse, semaphore execution) {
	defer semaphore.release()
	var count int64

	for jobInput := range h.selfQueue {
		jobOutputChan <- businessFunctionalityJob(jobInput)
		count++
	}
	fmt.Println("我执行了", count, "个任务")
}

func TryConsistWorkerPool() {
	numberOfJobs := 100
	// 任务channel
	jobsChan := make(chan JobInput, numberOfJobs)
	// 结果channel
	resultsChan := make(chan JobResponse, numberOfJobs)

	// worker pool init  生产结果定向到 resultsChan
	numOfWorkers := 10
	hardWorkerSemaphore := NewExecutionSlots(numOfWorkers)
	defer hardWorkerSemaphore.close()

	workerIndex := make(map[int]*hardWorker)
	for i := 0; i < numOfWorkers; i++ {
		hardWorkerSemaphore.acquire()
		worker := NewHardWorker(numberOfJobs)
		workerIndex[i] = worker
		go worker.doWorkUsingSelfQueue(resultsChan, hardWorkerSemaphore)
	}

	// 任务拆解 生产
	go jobProduce100(jobsChan)
	//go jobProduce100Static(jobsChan)
	// 任务分发 分发
	go jobDispatch(jobsChan, workerIndex)

	// 结果整合
	var responses []JobResponse
	collectWorkerSemaphore := NewExecutionSlots(1)
	defer collectWorkerSemaphore.close()
	collectWorkerSemaphore.acquire()
	go func() {
		defer collectWorkerSemaphore.release()
		// resultsChan 关闭时结束
		for resp := range resultsChan {
			responses = append(responses, resp)
		}
	}()

	// 等待所有work工作完毕
	hardWorkerSemaphore.wait()
	// 前面结果已经全部推送到 resultsChan 可以关闭了
	close(resultsChan)

	// 结果回收
	collectWorkerSemaphore.wait()
	// 结果整理
	fmt.Println(combineResponses(responses).finalOutput)
}

func jobDispatch(jobsChan chan JobInput, index map[int]*hardWorker) {
	count := len(index)
	for jobInput := range jobsChan {
		id := jobInput.startTime.Unix() % int64(count)
		fmt.Println("task 分发给", id)
		index[int(id)].selfQueue <- jobInput
	}
	// 分发结束 关闭worker队列
	for _, v := range index {
		close(v.selfQueue)
	}
	fmt.Println("Dispatch finish.")
}
