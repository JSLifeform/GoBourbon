package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

const (
	IntegerCol = "INT"
	NumericCol = "NUMERIC(6,2)"
	StringCol  = "VARCHAR(255)"
	FilePath   = `Bourbon_Website_Data.xlsx`
	Database   = `testdb`
)

type queryResult struct {
	ID            int     `json:"id"`
	Bourbon       string  `json:"bourbon"`
	Distillery    string  `json:"distillery"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Age           int     `json:"age"`
	Proof         float32 `json:"proof"`
	Estimated_Age string  `json:"estimated_age"`
	Aroma_Notes   string  `json:"aroma_notes"`
	Flavor_Notes  string  `json:"flavor_notes"`
	Finish_Notes  string  `json:"finish_notes"`
	Mashbill      string  `json:"mashbill"`
	Sourced       string  `json:"sourced"`
	Sourced_From  string  `json:"sourced_from"`
	Released      string  `json:"released"`
}

func dropTable(dbName string) string {
	dropSyntax := "DROP TABLE IF EXISTS " + dbName + ";"
	return dropSyntax
}

// function to select proper column type for SQL table creation statement
func columnType(columnName string) string {
	var input string
	// ask for and reads in data type
	fmt.Print("What data type(int, str, num) do you want for column ", columnName, "?")
	// reads in user input from CLI
	fmt.Scanln(&input)
	switch input {
	case "str":
		return StringCol
	case "int":
		return IntegerCol
	case "num":
		return NumericCol
	default:
		fmt.Println("Incorrect input!")
		return ""
	}
}

func main() {

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bourbon")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Panic on database open")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

	// open Excel file
	f, err := excelize.OpenFile(FilePath)
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

	//drops old table
	dropResults, err := db.Query(dropTable(Database))
	if err != nil {
		fmt.Println("Panic during table drop!")
		panic(err.Error())
	}
	defer dropResults.Close()

	// loops through rows in sheet
	for i, row := range rows {

		// skips first row of senseless data
		if i == 0 {
			continue

			// uses second row to create database with names of columns
		} else if i == 1 {

			var inputs []string
			for _, ColCell := range row {
				// loop to check for proper column type
				colType := ""
				for {
					colType = columnType(ColCell)
					if colType == "" {
						continue
					} else {
						break
					}
				}

				addToInput := ColCell + " " + colType + ", "
				inputs = append(inputs, addToInput)

				// if input == "str" {
				// 	addToInput := colCell + " " + StringCol + ", "
				// 	inputs = append(inputs, addToInput)
				// } else if input == "int" {
				// 	addToInput := colCell + " " + IntegerCol + ", "
				// 	inputs = append(inputs, addToInput)
				// } else if input == "num" {
				// 	addToInput := colCell + " " + NumericCol + ", "
				// 	inputs = append(inputs, addToInput)
				// } else {

				// 	// quits out of program, would like to prompt user again for correct input
				// 	fmt.Println("Incorrect input given, quitting until I figure out how to restart the loop and get correct input")
				// 	return
				// }

			}
			tableCreate := "CREATE TABLE IF NOT EXISTS " + Database + "(\n" + "id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, "
			for _, item := range inputs {
				tableCreate += item
			}
			// trims trailing ", " from tableCreate string
			tableCreate = strings.Trim(tableCreate, ", ")

			// appends trailing ");" to SQL query
			tableCreate += ");"

			// attempts to create table with SQL query
			fmt.Print(tableCreate, "\n\n\n")
			insert, err := db.Query(tableCreate)
			if err != nil {
				fmt.Println("Error on table create!")
				panic(err.Error())
			}

			// close insert
			defer insert.Close()

		} else {
			// this block does individual insert statements

			// string variable to add INSERT statement to
			addTable := "INSERT INTO " + Database + " VALUES ( NULL, "

			// loops through cells in row
			for _, colCell := range row {
				fmt.Println(colCell)

				// if colCell is empty add NULL statement to addTable query, otherwise add colCell information
				if colCell == "" {
					addTable += "NULL, "
				} else {
					addTable += `"` + colCell + `", `
				}

			}

			// trims trailing ', ' from string before adding closing line to statement
			addTable = strings.Trim(addTable, ", ")
			// adds trailing parentheses and semicolon to INSERT statement
			addTable += ");\n"

			// prints result and attempts insert query
			fmt.Println("Result is: ", addTable)
			insert, err := db.Query(addTable)
			if err != nil {
				fmt.Print("Error on ", addTable, "\n\n")
				fmt.Println(err)
				panic(err.Error())
			}

			// close insert statement
			defer insert.Close()
		}
	}

	results, err := db.Query("SELECT * FROM " + Database)
	if err != nil {
		fmt.Println("Panic on select statement")
		panic(err.Error())
	}

	//close query
	defer results.Close()

	for results.Next() {

		var PrintQuery queryResult
		err = results.Scan(&PrintQuery.ID, &PrintQuery.Bourbon, &PrintQuery.Distillery, &PrintQuery.City, &PrintQuery.State, &PrintQuery.Age, &PrintQuery.Proof,
			&PrintQuery.Estimated_Age, &PrintQuery.Aroma_Notes, &PrintQuery.Flavor_Notes, &PrintQuery.Finish_Notes, &PrintQuery.Mashbill, &PrintQuery.Sourced,
			&PrintQuery.Sourced_From, &PrintQuery.Released)
		if err != nil {
			fmt.Println("Panic on Scan")
			panic(err.Error())
		}
		fmt.Println(PrintQuery)
	}

}
