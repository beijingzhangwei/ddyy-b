package tools

import (
	"fmt"
)

func startWorkUsingSemaphore(jobInputChan <-chan JobInput, jobOutputChan chan<- JobResponse, semaphore execution) {
	defer semaphore.release()
	// jobInputChan 关闭时结束
	for jobInput := range jobInputChan {
		jobOutputChan <- businessFunctionalityJob(jobInput)
	}
}

func TrySemaphore() {
	// 任务channel
	jobsChan := make(chan JobInput, 10)

	// 结果channel
	resultsChan := make(chan JobResponse, 10)

	// worker pool init  生产结果定向到 resultsChan
	numOfWorkers := 3
	hardWorkerSemaphore := NewExecutionSlots(numOfWorkers)
	defer hardWorkerSemaphore.close()

	for i := 0; i < numOfWorkers; i++ {
		hardWorkerSemaphore.acquire()
		go startWorkUsingSemaphore(jobsChan, resultsChan, hardWorkerSemaphore)
	}

	// 任务拆解 生产
	go jobProduce(jobsChan)

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
