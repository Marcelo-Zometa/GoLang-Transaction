package main

import (
	"database/sql"
	"fmt"
	"log"
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

//Gets the user input for what to do in Menu
func getInt() int {
	var num int
	fmt.Scan(&num)	
	return num
}

func getFloat() float64 {
	var num float64

	fmt.Scan(&num)	
	num = math.Floor(num * 100) / 100

	return num
}

func getString() string {
	reader := bufio.NewReader(os.Stdin)
	//var mystr string
	mystr, _ := reader.ReadString('\n')
	return mystr
}

func menu(choice int) bool {
	
	switch choice {
		case 1:	//Get total incomes
				//caseOne()
		case 2: //Add an income
				caseTwo()
				break
		case 3: //Get total outcomes

		case 4: // Add an outcome

		case 5: //Exit the program
				return false
	}

	return true
}

// func caseOne() {
// 	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
// 		if errdb != nil {
// 			fmt.Println(" Error open db:", errdb.Error())
// 		}
// }

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
	if errdb != nil {
		log.Fatal(errdb)
	  }
	  fmt.Println("New record ID is:", id) 

	defer condb.Close()
}

//Driver function
func main() {
	var choice int
	//var runBack bool
	
	displayMenu()
	choice = getInt()
	
	menu(choice)
	//runBack = menu(choice)
}