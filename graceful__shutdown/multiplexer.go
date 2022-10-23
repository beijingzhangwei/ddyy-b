package main

import (
	"context"
	"log"
	"net/http"
)

// 我们需要做两件事: 分解如何注册路由和启动服务器。
//   1 幸运的是，我们可以拦截这个信号，以防止应用程序突然停止
//
//	// 2 服务器服务于 http 包，我们无法控制停止它。幸运的是，我们
//	//   可以简单地创建自己的 http 服务器，这样就可以更好地控制如何停止它。
//
//	// 3 当我们需要关闭服务器时，请求可能正在处理，我们如何等待请求先完成？
//	//   幸运的是，由于我们有自己的 http 服务器，有一个方便的方法可以优雅地关闭服务器！
//
//	// 4 当服务器正在等待请求完成处理时，其他请求可能仍在流中。像上面的问题一样，
//	//   方便的关机方法也可以处理这个问题！
func multiplexer() {
	// (1) create a server mux (mux stands for Multiplexer)
	//  首先，我们需要创建一个多路复用器(简称 mux) ，它基本上是一个请求处理程序
	serverMux := http.NewServeMux()

	// (2) register route and handler
	// 使用多路复用器，我们注册路由和处理程序。与我们之前使用 http 包完成的操作非常相似
	serverMux.HandleFunc("/hello", HelloAPI)

	// (3) create the server
	//     然后我们创建服务器。有了这个对象，以后我们将了解关闭方法
	server := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	// (4) start the server
	//  然后我们通过调用... ListenAndServe 启动服务器
	err := server.ListenAndServe()

	// (5) handle error necessarily on exit
	//     如果服务器无法启动，我们可能应该记录错误
	if err != nil {
		// (6) if server exits with ErrServerClosed, it meant graceful shutdown which we can ignore
		//      我们需要检查错误是否是由优雅的关机造成的，并安全地忽略它
		if err == http.ErrServerClosed {
			// do nothing...
		} else {
			log.Println(err)
		}
	}
	//
	//...
	//...

	// (7) shutdown the server gracefully using this method
	//     我们需要以某种方式调用 Shutdown 方法以正确拒绝请求的方式优雅地关闭服务器，并等待任何处理请求完成。
	err = server.Shutdown(context.Background())

	// (8) log the error on graceful exit
	//     最后，如果不能正确执行关机，我们也会记录错误
	if err != nil {
		panic(err)
	}
}
