package main

import "fmt"

func displayMenu() {
	fmt.Println("Greetings Marcelo\n")
	fmt.Println("Welcome back to your Expense Tracker\n")
	fmt.Println("What is it going to be for today?")
	fmt.Println("[1]Get total incomes")
	fmt.Println("[2]Add an income")
	fmt.Println("[3]Get total outcomes")
	fmt.Println("[4]Add an outcome")
}

func main() {
	displayMenu()
}