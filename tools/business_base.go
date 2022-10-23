package tools

import (
	"fmt"
	"time"
)

type JobInput struct {
	startTime time.Time
	endTime   time.Time
}

type JobResponse struct {
	//keeping this empty for this example. Real-world will be more complex.
}

type Response struct {
	finalOutput string
}

func combineResponses(jobResponses []JobResponse) Response {
	return Response{finalOutput: "Well done!!"}
}
func businessFunctionalityJob(jobInput JobInput) JobResponse {
	fmt.Println("Executing Job..")
	time.Sleep(10 * time.Millisecond)
	return JobResponse{}
}
func jobProduce(jobInputChan chan<- JobInput) {
	jobInputs := []JobInput{JobInput{startTime: time.Now().Add(time.Second)}, JobInput{startTime: time.Now()}}
	for _, jobInput := range jobInputs {
		jobInputChan <- jobInput
	}
	// 任务输入完毕 关闭通道（ : range jobInputChan 随之结束）
	close(jobInputChan)
	fmt.Println("JobInputs finish.")
}

func jobProduce100(jobInputChan chan<- JobInput) {
	for i := 0; i < 100; i++ {
		jobInputChan <- JobInput{startTime: time.Now().Add(time.Duration(i) * time.Second)}
	}
	// 任务输入完毕 关闭通道（ : range jobInputChan 随之结束）
	close(jobInputChan)
	fmt.Println("JobInputs finish.")
}

func jobProduce100Static(jobInputChan chan<- JobInput) {
	for i := 0; i < 100; i++ {
		jobInputChan <- JobInput{startTime: time.Now().Add(1 * time.Second)}
	}
	// 任务输入完毕 关闭通道（ : range jobInputChan 随之结束）
	close(jobInputChan)
	fmt.Println("JobInputs finish.")
}
