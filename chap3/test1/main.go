package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Customer struct {
	id        int
	name      string
	address   string
	contactno string
}

var customers []Customer

func authenicate() bool {
	var userid, password string
	fmt.Println("Enter your userid: ")
	fmt.Scan(&userid)
	fmt.Println("Enter your password: ")
	fmt.Scan(&password)

	if userid == "guest" && password == "guest" {
		fmt.Println("You are authenticated.")
		return true
	} else {
		fmt.Println("Your input is wrong. Please enter again")
		return false
	}
}

func showCommandMenu() {
	fmt.Println("1. Display Customers")
	fmt.Println("2. Display Customer by ID")
	fmt.Println("3. Enter Customer Information")
	fmt.Println("4. Update Customer Information")
	fmt.Println("5. Delete Customer")
	fmt.Println("0. Exit")
	fmt.Print("Enter the menu choice id:")
}

func CommandMenu() {
	var choice int
	showCommandMenu()
	for {
		fmt.Scan(&choice)
		if choice == 0 {
			fmt.Println("You entered 0 for exiting the program.")
			break
		} else {
			switch choice {
			case 1:
				displayCustomerList(customers)
				fmt.Println("Input 'Enter' to go back to commands menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 2:
				var customerId int
				fmt.Println()
				fmt.Print("Input Customer ID :")
				fmt.Scan(&customerId)
				fmt.Println("\nDisplaying Customer by ID ", customerId)
				displayCustomerByID(customerId, customers)
				fmt.Println("Input 'Enter' to go back to commands menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 3:
				fmt.Println("Add a new Customer")
				addNewCustomer(&customers)
				fmt.Println()
				fmt.Println("Input 'Enter' to go back to commands menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 4:
				var customerId int
				fmt.Print("Input Customer Id: ")
				fmt.Scan(&customerId)
				updateCustomerById(customerId, &customers)
				fmt.Println()
				fmt.Println("Input 'Enter' to go back to commands menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 5:
				var customerId int
				fmt.Println()
				fmt.Print("Input Customer Id: ")
				fmt.Scan(&customerId)
				deleteCustomerById(customerId, &customers)
				fmt.Println()
				fmt.Println("Input 'Enter' to go back to commands menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			default:
				fmt.Println("Input your choice again.")
				showCommandMenu()
			}
		}
	}
}

func displayCustomerList(customers []Customer) {
	for _, customer := range customers {
		fmt.Println("ID         :", customer.id)
		fmt.Println("Name       :", customer.name)
		fmt.Println("Address    :", customer.address)
		fmt.Println("Contact No :", customer.contactno)
		fmt.Println()
	}
}

func displayCustomerByID(id int, customers []Customer) {
	for _, customer := range customers {
		if customer.id == id {
			fmt.Println("ID         :", customer.id)
			fmt.Println("Name       :", customer.name)
			fmt.Println("Address    :", customer.address)
			fmt.Println("Contact No :", customer.contactno)
			fmt.Println()
			break
		}
	}
}

func addNewCustomer(customers *[]Customer) {
	var customerId int
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter Customer Id : ")
	fmt.Scan(&customerId)
	fmt.Print("Enter Customer Name : ")
	name, _ := consoleReader.ReadString('\n')
	fmt.Print("Enter Customer address : ")
	address, _ := consoleReader.ReadString('\n')
	fmt.Print("Enter Customer contactno : ")
	contactno, _ := consoleReader.ReadString('\n')

	Customer_ := Customer{
		id:        customerId,
		name:      strings.TrimSpace(name),
		address:   strings.TrimSpace(address),
		contactno: strings.TrimSpace(contactno),
	}

	*customers = append(*customers, Customer_)
}

func updateCustomerById(id int, customers *[]Customer) {
	var newContactno string
	fmt.Println()
	fmt.Print("Enter the new contact number of the Customer: ")
	fmt.Scan(&newContactno)
	for i, customer := range *customers {
		if customer.id == id {
			(*customers)[i] = Customer{
				id,
				customer.name,
				customer.address,
				newContactno,
			}
		}
	}
}

func deleteCustomerById(id int, customers *[]Customer) {
	for i, customer := range *customers {
		if customer.id == id {
			*customers = append((*customers)[:i], (*customers)[i+1:]...)
		}
	}
}

func CRMConsoleApp() {
	for !authenicate() {
		authenicate()
		break
	}

	customers = append(customers, Customer{
		id:        1,
		name:      "Jhon Smith",
		address:   "California",
		contactno: "231456",
	})

	customers = append(customers, Customer{
		id:        2,
		name:      "Brad Smith",
		address:   "Boston",
		contactno: "4321234",
	})

	CommandMenu()
}

func main() {
	CRMConsoleApp()
}
