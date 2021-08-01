package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}

	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		return srv.Shutdown(ctx)
	})
	wg.Add(1)
	eg.Go(func() error {
		wg.Done()
		return srv.ListenAndServe()
	})


	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)

	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sigs:
				signal.Stop(sigs)
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}

}

func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}
