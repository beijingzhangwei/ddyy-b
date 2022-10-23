package tools

import (
	"fmt"
	"sync"
)

func startWorker(jobInputChan <-chan JobInput, jobOutputChan chan<- JobResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	// jobInputChan 关闭时结束
	for jobInput := range jobInputChan {
		jobOutputChan <- businessFunctionalityJob(jobInput)
	}
}

func Try() {
	// 任务channel
	jobsChan := make(chan JobInput, 10)

	// 结果channel
	resultsChan := make(chan JobResponse, 10)

	// worker pool init  生产结果定向到 resultsChan
	numOfWorkers := 3
	wg := sync.WaitGroup{}
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go startWorker(jobsChan, resultsChan, &wg)
	}

	// 任务拆解 生产
	go jobProduce(jobsChan)

	// 结果整合
	var responses []JobResponse
	wgResp := sync.WaitGroup{}
	wgResp.Add(1)
	go func() {
		defer wgResp.Done()
		// resultsChan 关闭时结束
		for resp := range resultsChan {
			responses = append(responses, resp)
		}
	}()

	// 等待所有work工作完毕
	wg.Wait()
	// 前面结果已经全部推送到 resultsChan 可以关闭了
	close(resultsChan)

	// 结果回收
	wgResp.Wait()
	// 结果整理
	fmt.Println(combineResponses(responses).finalOutput)
}
