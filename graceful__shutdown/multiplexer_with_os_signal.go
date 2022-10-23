package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// print that the application has started
	fmt.Println("Hello, Graceful Server!")

	// create a server mux
	serverMux := http.NewServeMux()

	// register route and handler
	serverMux.HandleFunc("/hello", HelloAPI_)

	// create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	// create channel to capture signal
	sigChn := make(chan os.Signal, 1)

	// register channel for signal capturing
	signal.Notify(sigChn, syscall.SIGINT)

	// create channel to capture server error
	startErrChn := make(chan error, 1)

	// start the server asynchronously
	go func() {
		// (4) start the server
		go func() {
			time.Sleep(10 * time.Second)
			err := server.Shutdown(context.Background()) // 这里主动关停，会直接中断处理中的请求
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		err := server.ListenAndServe()

		// (5) handle error on exit
		if err != nil {
			if err == http.ErrServerClosed {
				// do nothing...
			} else {
				// log error
				fmt.Println(err)
			}
		}

		// inform that server has stopped accepting requests
		startErrChn <- err
	}()

	// wait for either a Ctrl+C signal, or server abnormal start error
	select {

	// we captured a signal to shut down application
	case <-sigChn:
		// print that server is shutting down
		fmt.Println("server is shutting down")

		// trigger the server shutdown gracefully
		err := server.Shutdown(context.Background())

		// log any error on graceful exit
		if err != nil {
			fmt.Println(err)
		}

	// we have an error from server's listen and serve, which is abnormal shutdown
	case <-startErrChn:
		// since we already logged the error, we may want to log additional details
		fmt.Println("server abnormal shutdown without stop signal!")
	}

	// tidy up and print that we have gracefully shutdown the server
	fmt.Println("Graceful shutdown!")
}

// a long running request
func HelloAPI_(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Second)
	response := "Hello!"
	w.Write([]byte(response))
}
