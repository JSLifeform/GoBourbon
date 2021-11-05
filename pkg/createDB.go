package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	integerCol = "int"
	numericCol = "numeric(6,2)"
	stringCol  = "varchar(45)"
	filePath   = `C:\Users\Owner\go\bourbon_database\Bourbon_Website_Data.xlsx`
	database   = `testdb`
)

func main() {
	//columnlist := []string("")

	// open Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// read in rows individually
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// loops through rows in sheet
	for _, row := range rows {

		// prints insert part of SQL statement
		fmt.Println("INSERT INTO", database, "VALUES (")

		// loops through cells in row
		for _, colCell := range row {
			fmt.Print(`"`, colCell, `"`, ", ")
		}
		// I think adds new line, delete?
		fmt.Println(");\n")
	}
	// shows opened file
	fmt.Println(filePath)

}
