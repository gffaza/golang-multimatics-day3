package pustaka
import(
	"fmt"
	"time"
)

func JalanAntrianChannel(){
	ch := make(chan string)
	
	go sendMessage("Go routine", ch)

	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println("Pengiriman pesan selesai!")
}

func sendMessage(message string, ch chan string){
	for i := 1; i<= 5; i++{
		ch <- fmt.Sprintf("%s, %d", message, i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}