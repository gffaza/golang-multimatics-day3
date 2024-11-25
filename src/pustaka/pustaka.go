package pustaka
import "fmt"

func BilangHalo(){
	fmt.Println("Hai")
}

var Anon = func(angka ...int)(int, float64){
	total := 0
	for _, v := range angka {
		total +=v
	}


	rata2 := float64(total)/ float64(len(angka))
	return total, rata2
}
