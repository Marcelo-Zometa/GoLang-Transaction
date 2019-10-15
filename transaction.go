package main

import (
	"database/sql"
	"fmt"
	//"log"
	_ "github.com/denisenkom/go-mssqldb"
	"time"
	"math"
	"os"
	"bufio"
  )
//Displays menu options. 
func displayMenu() {
	fmt.Println("Greetings Marcelo\n")
	fmt.Println("Welcome back to your Expense Tracker\n")
	fmt.Println("What is it going to be for today?")
	fmt.Println("[1]Get total incomes")
	fmt.Println("[2]Add an income")
	fmt.Println("[3]Get total outcomes")
	fmt.Println("[4]Add an outcome")
	fmt.Println("[5]Exit the program")
}

//Gets the user input of int
func getInt() int {
	var num int
	fmt.Scan(&num)	
	return num
}
//Gets the user input of float64
func getFloat() float64 {
	var num float64

	fmt.Scan(&num)	
	num = math.Floor(num * 100) / 100

	return num
}
//Gets the user input of string
func getString() string {
	reader := bufio.NewReader(os.Stdin)
	//var mystr string
	mystr, _ := reader.ReadString('\n')
	return mystr
}

func menu(choice int) {
	
	switch choice {
		case 1:	//Get total incomes
				//caseOne()
		case 2: //Add an income
				caseTwo()
				break
		case 3: //Get total outcomes

		case 4: // Add an outcome
				caseFour()
				break
		case 5: //Exit the program
				fmt.Println("Thanks for using your Expense Tracker. Come again soon!")
				os.Exit(0)
	}
}

// func caseOne() {
// 	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
// 		if errdb != nil {
// 			fmt.Println(" Error open db:", errdb.Error())
// 		}
// }

//Displays the total incomes from the database.
func caseOne() {
	
}

//Gets the input from user of description of income, amount and injects it into the Income database
func caseTwo() {
	//Creation of variables
	var sqlStatement string
	var money float64
	var income string
	id := 0

	//Getting the info from user
	fmt.Println("Input the name of this income: ")
	income = getString()

	fmt.Println("Input the amount of this income: ")
	money = getFloat()

	//Creating connection
	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		}
		
	//Preparing query and executing it
	sqlStatement = "Insert into F19MarceloZ.GoLangTransactions.Income (income, dateInput, amount) VALUES ($1, $2, $3);"

	errdb = condb.QueryRow(sqlStatement, income, time.Now(), money).Scan(&id)
	// if errdb != nil {
	// 	log.Fatal(errdb)
	//   }
	  fmt.Println("Entry successfully added to database") 

	defer condb.Close()
}

//Gets the input from user of description of outcome, amount and injects it into the Outcome database
func caseFour() {
	//Creation of variables
	var sqlStatement string
	var money float64
	var outcome string
	id := 0

	//Getting the info from user
	fmt.Println("Input the name of this outcome: ")
	outcome = getString()

	fmt.Println("Input the amount of this outcome: ")
	money = getFloat()

	//Creating connection
	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		}
		
	//Preparing query and executing it
	sqlStatement = "Insert into F19MarceloZ.GoLangTransactions.Outcome (outcome, dateInput, amount) VALUES ($1, $2, $3);"

	errdb = condb.QueryRow(sqlStatement, outcome, time.Now(), money).Scan(&id)
	// if errdb != nil {
	// 	log.Fatal(errdb)
	//   }
	  fmt.Println("Entry successfully added to database") 

	defer condb.Close()
}

//Driver function
func main() {
	var choice int
	
	for {
		displayMenu()
		choice = getInt()
		
		menu(choice)
		//runBack = menu(choice)
	}	
}