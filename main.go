package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	results, err := db.Query("SELECT id, testcol FROM test")
	if err != nil {
		fmt.Println("Panic on select statement")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			fmt.Println("Panic on Scan")
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag.Name)
	}

}
