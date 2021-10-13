package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		go func() {
			i := i
			ch <- i
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		val := <-ch
		println(val)
		if i == 2 {
			close(ch)
			time.Sleep(time.Second * 2)
			return
		}
	}
}
