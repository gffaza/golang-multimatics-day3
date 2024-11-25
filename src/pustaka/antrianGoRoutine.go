package pustaka
import (
	"fmt"
	"time"
)

func JalanAntrian(){

	// awalan go pasti go routine (asyncronus)
	go printMessage("Goroutine 1 ") 
	go printMessage("Goroutine 2")

	printMessage("Main Goroutine")

}

func printMessage(message string){
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(1 * time.Second)		
	}
}