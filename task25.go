package main

import (
	"fmt"
	"sync"
	"time"
)

// 下記関数を同時に実行
func printHello() {
	fmt.Println("Hello")
}

func printWorld() {
	fmt.Println("World")
}

func task25_1() {
	fmt.Println("① Start")
	go printHello()
	go printWorld()

	// routinesが完了できるように休止 = タスク完了を待つわけではないため実装では使用しない
	time.Sleep(1 * time.Second)
	fmt.Println("① Done")
}

func task25_2() {
	fmt.Println("② Start")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Hello from go routine")
		}()
	}
	// 全てのroutineが完了するまで待機
	wg.Wait()
	fmt.Println("② Done")
}
