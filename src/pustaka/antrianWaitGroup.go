package pustaka

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func JalanAntrianWG(){
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	wg.Add(3)

	go printMessageWG("Goroutine 1", &wg)
	go printMessageWG("Goroutine 2", &wg)
	go printMessageWG("Goroutine 3", &wg)

	wg.Wait()

	fmt.Println("Semua goroutine selesai!")
}

func printMessageWG(message string, wg *sync.WaitGroup){
	defer wg.Done()

	for	i := 1; i<=5; i++ {
		fmt.Println(message,i)
		time.Sleep(1*time.Second)
	}
}