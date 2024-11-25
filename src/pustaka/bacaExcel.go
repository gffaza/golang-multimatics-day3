package pustaka

import (
    "fmt"
    "time"
)

func Baca() {
	var wg sync.WaitGroup
	var waktuMulai = time.Now()
		counter := 0
		totalAmount := 0

		type RowData struct {
			index int
			Row *xlsx.Row
		}
		rowsChannel := make(chan RowData)

		filepath := "forTraining.xlsx"
		xlFile, err := xlsx.OpenFile(filepath)
		if err!= nill{
			log.Fatal("Couldn't open file: %v", err)
		}

		sheet := xlFile.NewSheet[0]

		results := make([]string, len(sheet.Rows))

		wg.Add(1)
		go func(){
			defer wg.Done()
			for i, row := range sheet.Rows{
				if i == 0{
					continue
				}
				rowsChannel <- RowData{Index : i, Row: row}
			}	
			close(rowsChannel)
		}()

		wg.Add(1)
		go func(){
			defer wg.Done()
            for rowData := range rowsChannel{

				id := rowData.Row.Cells[1].String()
				initiatorRefNo := rowData.Row.Cells[3].String()
				sysRefNo := rowData.Row.Cells[4].String()

				var amount string = "0"


				results[rowData.index] = fmt.Sprintf("ID: %s, Initiator: %s, SYS_REF_NO: %s, AMOUNT: %s", id, initiatorRefNo, sysRefNo, amount)
				convertAmount, errr := strconv.Atoi(amount)
				if errr!= nil{
					totalAmount, err += convertAmount
                }
            }
		}
}
