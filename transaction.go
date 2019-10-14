package main

import "fmt"

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
func getChoice() int {
	var num int

	return fmt.Scan(&num)	
}

func menu(choice int) bool {
	
	switch choie {
	case 1:	//Get total incomes

	case 2: //Add an income

	case 3: //Get total outcomes

	case 4: // Add an outcome

	case 5: //Exit the program
	}
}

//Driver function
func main() {
	var choice int
	var runBack bool

	displayMenu()
	choice = getChoice()
	
	runBack = menu(choice)
}