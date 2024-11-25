package main

import (
	"fmt"
)


func main() {
	fmt.Println("Go World!") 
	
	dHms := struct{
		nama string
		umur int
	}{
		"Budi",
		21,
	}

	dHms.umur = 50
	dHms.nama = "Anto"

	fmt.Println(dHms)

	type dKaryawan struct{
		nama string
		umur int
		alamat string
	}

	var arrData = []dKaryawan{
		{
			"Budi", 
			21, 
			"Jakarta",
		},
        {	
			"Anto", 
			50, 
			"Bandung",
		},
		{
			nama : "Sinta",
			umur : 22,
            alamat : "Surabaya",
		},
	}

	fmt.Println(arrData)
	for index, isi := range arrData {
		fmt.Printf("Data Karyawan %d: %v\n", index+1, isi)
	}



}