package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			if !Status() {
				Alert()
				fmt.Println("error")
			} else {
				fmt.Println("ok")
			}
			// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1
	}()
	<-ch
}
