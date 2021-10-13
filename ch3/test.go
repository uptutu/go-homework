package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch := make(chan int, 0)
	for i := 0; i < 10; i++ {
		go send(ch)
	}

	time.Sleep(time.Second * 3)
	recv(ch)
	time.Sleep(time.Second * 3)

}

func send(ch chan int) {
	ch <- 1
}

func recv(ch chan int) {
	fmt.Println(<-ch)
}
