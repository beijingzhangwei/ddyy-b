package main

import (
	"fmt"
	"net/http"
)

func main_() {
	fmt.Println("Hello, Graceful Server!")

	// (1) create channel to receive the signal
	// (1)创建一个缓冲区刚好能够捕获 Ctrl + C 信号的通道
	//sigChn := make(chan os.Signal, 1)

	// (2) let Go inform us when signal is captured
	//     当信号被拦截时，我们让 Go 系统通知我们
	//signal.Notify(sigChn, syscall.SIGINT)

	// (3) now we run our server in the background
	//      由于我们需要等待主线程上的信号并在它被调用时处理它，
	//      我们应该移动我们的 http。ListenAndServe 代码编写成一个 goroutine
	//go func() {
	//	//...listen and serve here...
	//}()

	// (4) wait for signal, this will pause the main application here
	//     we pause the main thread and wait for the signal
	//     暂停主线程，等待信号
	//     只有3行简单的代码允许我们捕获 Ctrl + C 信号。
	//     通过这样做，我们解决了第一个突然退出应用程序的问题。
	//<-sigChn

	// (1) api route and handler
	http.HandleFunc("/hello", HelloAPI)

	// (2) start the server
	err := http.ListenAndServe(":8080", nil)

	// (3) handle error on exit
	if err != nil {
		panic(err)
	}
}

// (4) A simple basic implementation of handling an API request
func HelloAPI(w http.ResponseWriter, r *http.Request) {
	response := "Hello!"
	w.Write([]byte(response))
}
