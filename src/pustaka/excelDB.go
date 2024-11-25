package pustaka

import (
	"fmt"
	"log"
	"main/koneksi"
	"os"
	"sync"
	"time"

	"github.com/tealeg/xlsx"
) 

func TulisDB(){
	var wg sync.WaitGroup

	type RowData struct {
		Index int 
		Row *xlsx.Row
	}
	rowsChannel := make(chan RowData)

	filePath :="forTraining.xlsx"
	xlFile, err := xlsx.OpenFile(filePath)
	if err!= nil{
		log.Fatalf("Error opening file: %v", err)
	}

	sheet := xlFile.Sheets[0]

	db, _:= koneksi.Konek()

	go func(){
		defer wg.Done()
		for i , row := range sheet.Rows {
			if i == 0 {
				continue
			}
			rowsChannel <- RowData{Index: i, Row: row}
		}
		close(rowsChannel)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		defer db.Close()
		for rowData := range rowsChannel {
			id := rowData.Row.Cells[1].String()
			initiatorRefNo := rowData.Row.Cells[3].String()
			sysRefNo := rowData.Row.Cells[4].String()
			amount := rowData.Row.Cells[12].String()

			insertQuerry := "INSERT INTO ngajargolang (ID, INITIATOR_REF_NO, SYS_REF_NO, amount) Values (?, ?, ?, ?)"
			_, err := db.Exec(insertQuerry, id, initiatorRefNo, sysRefNo, amount)
			if err != nil{
				log.Fatalf("error insert data (ID %s): %v, id, err")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Data successfully inserter into MySQL.")
}

func BacaDB() error {
	db, _ := koneksi.Konek()
	rows, err := db.Query("SELECT ID, INITIATOR_REF_NO, SYS_REF_NO FROM ngajargolang")
	if err != nil {
		return fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return fmt.Errorf("error creating sheet: %v", err)
	}

	header := sheet.AddRow()
	header.AddCell().Value = "ID"
	header.AddCell().Value = "INITIATOR_REF_NO"
	header.AddCell().Value = "SYS_REF_NO"

	for rows.Next() {
		var id, initiatorRefNo, sysRefNo string

		err := rows.Scan(&id, &initiatorRefNo, &sysRefNo)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		row := sheet.AddRow()
		row.AddCell().Value = id
		row.AddCell().Value = initiatorRefNo
		row.AddCell().Value = sysRefNo
	}

	err = file.Save("output.xlsx")
	if err != nil {
		return fmt.Errorf("error saving excel file: %v", err)
	}

	return nil

}

func Csv() error {
	t0 := time.Now()
	db, _ := koneksi.Konek()
	rows, err := db.Query("SELECT ID, INITIATOR_REF_NO, SYS_REF_NO, AMOUNT FROM ngajargolang")
	if err != nil {
		log.Fatalf("Error querying database: %s", err)
	}
	defer rows.Close()

	file, err := os.Create("./forTrainingFromDB.csv")
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	// write header
	_, err = file.WriteString("ID,INITIATOR_REF_NO,SYS_REF_NO,AMOUNT\n")
	if err != nil {
		return fmt.Errorf("error writing header: %s", err)
	}

	// loop through the rows
	for rows.Next() {
		var id, initiatorRefNo, sysRefNo, amount string

		err := rows.Scan(&id, &initiatorRefNo, &sysRefNo, &amount)
		if err != nil {
			return fmt.Errorf("error scanning rows: %s", err)
		}

		_, err = file.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", id, initiatorRefNo, sysRefNo, amount))
		if err != nil {
			return fmt.Errorf("error writing rows: %s", err)
		}
	}

	err = file.Sync()
	if err != nil {
		return fmt.Errorf("error syncing file: %s", err)
	}

	log.Println("Data has been written to csv file")
	t1 := time.Now()

	log.Printf("The query took %v to run\n", t1.Sub(t0))

	return nil
}