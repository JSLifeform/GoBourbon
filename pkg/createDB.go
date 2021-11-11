package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	integerCol = "INT"
	numericCol = "NUMERIC(6,2)"
	stringCol  = "VARCHAR(45)"
	filePath   = `Bourbon_Website_Data.xlsx`
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
	for i, row := range rows {

		// skips first row of senseless data
		if i == 0 {
			continue

			// uses second row to create database with names of columns
		} else if i == 1 {

			var inputs []string
			for _, colCell := range row {
				// string for user to input column type
				var input string
				// ask for and reads in data type
				fmt.Println("What data type(int, str, num) do you want for column ", colCell, "?")
				fmt.Scanln(&input)
				if input == "str" {
					addToInput := colCell + " " + stringCol + ", "
					inputs = append(inputs, addToInput)
				} else if input == "int" {
					addToInput := colCell + " " + integerCol + ", "
					inputs = append(inputs, addToInput)
				} else if input == "num" {
					addToInput := colCell + " " + numericCol + ", "
					inputs = append(inputs, addToInput)
				} else {

					// quits out of program, would like to prompt user again for correct input
					fmt.Println("Incorrect input given, quitting until I figure out how to restart the loop and get correct input")
					return
				}

			}
			tableCreate := "CREATE TABLE [IF NOT EXISTS] " + database + "(\n" + "id INT AUTO_INCREMENT PRIMARY KEY, "
			for _, item := range inputs {
				tableCreate += item
			}
			// trims trailing ", " from tableCreate string
			tableCreate = strings.Trim(tableCreate, ", ")

			// appends trailing ");" to SQL query
			tableCreate += ");"
			fmt.Print(tableCreate, "\n\n\n")
		} else {

			// string variable to add INSERT statement to
			addTable := "INSERT INTO " + database + " VALUES ("

			// loops through cells in row
			for _, colCell := range row {
				addTable += `"` + colCell + `", `
			}

			// trims trailing ', ' from string before adding closing line to statement
			addTable = strings.Trim(addTable, ", ")
			// adds trailing parentheses and semicolon to INSERT statement
			addTable += ");\n"
			// I think adds new line, delete?
			fmt.Println(addTable)
		}
	}
	// shows opened file
	fmt.Println(filePath)

}

/*
package main

import (
	"fmt"
)

func getColType() string {
	for {
		var input string
		// ask for and reads in data type
		fmt.Println("What data type(int, str, num) do you want for column ", colCell, "?")
		fmt.Scanln(&input)
		switch input {
		case "str":
			return StrColType
		case "int":
			return IntColType
		case "num":
			return NumColType
		default:
			fmt.Println("Incorrect input given, quitting until I figure out how to restart the loop and get correct input")
		}
	}
}

func main() {
	// string for user to input column type

	addToInput := colCell + " " + colType + ", "
	inputs = append(inputs, addToInput)
}
*/
