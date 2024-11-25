package pustaka

import (
	"fmt"
	"sync"
	"time"
)

func JalanAntrianGabungan(){
	ch := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go sendMessageGabungan("Goroutine 1", ch, &wg)
	go sendMessageGabungan("Goroutine 2", ch, &wg)

	go func(){
		wg.Wait()
		close(ch)
	}()

	for msg := range ch{
		fmt.Println("Channel: " + msg)
	}

	fmt.Println("Semua pesan telah diterima danpengiriman selesai!")
}

func sendMessageGabungan(message string, ch chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for i := 1; i<= 5; i++{
		ch <- fmt.Sprintf("%s, %d", message, i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}