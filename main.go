package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
	// "Bourbon_Database/pkg/createDB"
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

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bourbon")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Panic on database open")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// // perform a db.Query insert
	// insert, err := db.Query("INSERT INTO test VALUES ( 6, 'HI EEVGENIYA' )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	fmt.Println("Panic on insert statement")
	// 	panic(err.Error())
	// }

	// // close insert statement
	// defer insert.Close()

	// test select statement

	// open Excel file
	f, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// rows, _ := openSpreadsheet(filePath)

	// read in rows individually
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	//drops old table
	dropTable := "DROP TABLE IF EXISTS " + Database + ";"
	DropResults, err := db.Query(dropTable)
	if err != nil {
		fmt.Println("Panic during table drop!")
		panic(err.Error())
	}
	fmt.Print("New table created! Results of db.Query: ", &DropResults, "\n\n\n")

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
					addToInput := colCell + " " + StringCol + ", "
					inputs = append(inputs, addToInput)
				} else if input == "int" {
					addToInput := colCell + " " + IntegerCol + ", "
					inputs = append(inputs, addToInput)
				} else if input == "num" {
					addToInput := colCell + " " + NumericCol + ", "
					inputs = append(inputs, addToInput)
				} else {

					// quits out of program, would like to prompt user again for correct input
					fmt.Println("Incorrect input given, quitting until I figure out how to restart the loop and get correct input")
					return
				}

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
		panic(err.Error()) // proper error handling instead of panic in your app
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
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(PrintQuery)

		// // var tag Tag
		// // for each row, scan the result into our tag composite object
		// err := results.Scan(&printStr)
		// if err != nil {
		// 	fmt.Println("Panic on Scan")
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }
		// // and then print out the tag's Name attribute
		// fmt.Println(printStr)
	}

}
