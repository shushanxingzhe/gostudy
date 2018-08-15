package main

import (
	"sync"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i:=1;i<10;i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("A:",i)
		}
	}()
	go func(){
		defer wg.Done()
		for i:=1;i<10;i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("B:",i)
		}
	}()
	wg.Wait()
	fmt.Println("finished")
}
