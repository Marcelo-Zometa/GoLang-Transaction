package main

import (
	"database/sql"
	"fmt"
	//"log"
	_ "github.com/denisenkom/go-mssqldb"
	"time"
  )
// //Displays menu options. 
// func displayMenu() {
// 	fmt.Println("Greetings Marcelo\n")
// 	fmt.Println("Welcome back to your Expense Tracker\n")
// 	fmt.Println("What is it going to be for today?")
// 	fmt.Println("[1]Get total incomes")
// 	fmt.Println("[2]Add an income")
// 	fmt.Println("[3]Get total outcomes")
// 	fmt.Println("[4]Add an outcome")
// 	fmt.Println("[5]Exit the program")
// }

// //Gets the user input for what to do in Menu
// func getChoice() int {
// 	var num int

// 	return fmt.Scan(&num)	
// }

// func menu(choice int) bool {
	
// 	switch choice {
// 		case 1:	//Get total incomes
				
// 		case 2: //Add an income

// 		case 3: //Get total outcomes

// 		case 4: // Add an outcome

// 		case 5: //Exit the program
// 				return false
// 	}

// 	return true
// }

//Driver function
func main() {
	//var choice int
	//var runBack bool
	var sqlStatement string

	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		}

	// displayMenu()
	// choice = getChoice()
	
	// runBack = menu(choice)
	
	sqlStatement = "Insert into F19MarceloZ.GoLangTransactions.Income (income, dateInput, amount) VALUES ($1, $2, $3);"

	id := 0
	money := 1000.26
	errdb = condb.QueryRow(sqlStatement, "Jugar Origins", time.Now(), money).Scan(&id)
	if errdb != nil {
		panic(errdb)
	  }
	  fmt.Println("New record ID is:", id) 
	
	//   var (
	// 	sqlversion string
	//   )
	//   rows, err := condb.Query("select @@version")
	//   if err != nil {
	// 	log.Fatal(err)
	//   }
	//   for rows.Next() {
	// 	err := rows.Scan(&sqlversion)
	// 	if err != nil {
	// 	  log.Fatal(err)
	// 	}
	// 	log.Println(sqlversion)
	//   }
	defer condb.Close()
	
}