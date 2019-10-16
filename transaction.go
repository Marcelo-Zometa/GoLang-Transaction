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
		case 1:	//Get your balance
				balance()
				break			
		case 2: //Get total incomes
				totalIncome()
				break
		case 3: //Add an income
				addIncome()
				break
		case 4: //Get total outcomes
				totalOutcome()
				break
		case 5: // Add an outcome
		 		caseFour()//Exit the program
				break
		case 6:	fmt.Println("Thanks for using your Expense Tracker. Come again soon!") //Exit the program
				os.Exit(0)
	}
}

//Subtracts the total income minus the total outcome
func balance() {
	// var _income float64
	// var _outcome float64

	_income := totalIncome()
	_outcome := totalOutcome()
	
	fmt.Println("Your current balance is ", (_income - _outcome))
}

//Displays the total incomes from the database.
func totalIncome() float64 {
	//Creating functions
	var total float64

	//Creating connection
	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		}

	//Querying the database
	rows, err := condb.Query("select sum(i.amount) from f19MarceloZ.GoLangTransactions.Income i")
	if err != nil{
		log.Fatal(err)
	} 
	defer rows.Close()

	//Getting data from database and printing it
	for rows.Next(){
		err := rows.Scan(&total)

		if err != nil{
			log.Fatal(err)
		}
		
		fmt.Println("Total income: ", total)
	}

	//Closing connection
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer condb.Close()

	return total
}

//Gets the input from user of description of income, amount and injects it into the Income database
func addIncome() {
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

//Displays the total outcome from the database.
func totalOutcome() float64 {
	//Creating functions
	var total float64

	//Creating connection
	condb, errdb := sql.Open("mssql", "server=snow-se-1.snow.edu;user id=F19MarceloZometa;password=Password1!;")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		}

	//Querying the database
	rows, err := condb.Query("select sum(i.amount) from f19MarceloZ.GoLangTransactions.Outcome i")
	if err != nil{
		log.Fatal(err)
	} 
	defer rows.Close()

	//Getting data from database and printing it
	for rows.Next(){
		err := rows.Scan(&total)

		if err != nil{
			log.Fatal(err)
		}
		
		fmt.Println("Total outcome: ", total)
	}

	//Closing connection
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer condb.Close()

	return total
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
	
	//WARNING: infinite loop. I only do this because you can exit the program in the menu.
	for {
		displayMenu()
		choice = getInt()		
		menu(choice)	
	}	
}