# GoBourbon


This project is the start of an interface between a SQL database and a future front end web page, which will contain a comprehensive list of bourbons in America. This project currently focuses on the database creation from an excel file, but will eventually contain more comprehensive SQL queries to search for bourbons via various characteristics.

GETTING STARTED

This project was written in Go version 1.17. It is recommended that anyone running it uses the same version. The SQL database was created using MySQL 8.0. Any SQL Managmenet tool should work with this project, but instructions will be written assuming the user is using MySQL 8.0. Go dependencies from github.com/go-sql-driver/mysql and github.com/xuri/excelize/v2 are required to run this project, but should be properly installed by the go mod init command in a future step.


INSTRUCTIONS

1.) Create the database "bourbon" within MySQL 8.0. For test purposes, giving access to this database with the simple root:password as the username/password combination, with local access at 127.0.0.1:3306. If any other credentials or access points are desired, the code can be changed on line 66. It is not necessary to populate any tables within this database, as the program will do it for you. 
2.) Open the Go terminal and navigate to the project folder.

3.) run the command "go mod init" to initialize the module and isntall the necessary Go dependencies.

4.) type "go run main.go" to run the project within its folder.

5.) The project will read all the columns from Excel file and prompt the user to type which kind of SQL column is required by the database. For testing purposes, the user can choose all columns to be str (for varchar(255) in SQL). The actual column requirements in future versions of this project are listed below:

    Bourbon str
    Distillery str
    City str
    State str
    Age int
    Proof num
    Estimated_Age str
    Aroma_Notes str
    Flavor_Notes str
    Finish_Notes str
    Mashbill str
    Sourced str
    Sourced_From str
    Released str

RESULTS

Currently, the project creates a table testdb and inserts the information from the Excel file into it. It then selects all the data in the table, places each row into a struct and displays it to the CLI. Future versions will allow for more complicated queries based on particular characters of the bourbons.

ACKNOWLEDGEMENTS

Special thank you to Code Louisville, our mentors Ron and Eric, and Pluralsight for allowing us the opportunity to dive into Go. The language seems simple yet powerful, and I look forward to taking the time after class to dive into this project further and expand my knowledge of Go.