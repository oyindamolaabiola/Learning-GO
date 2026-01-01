package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Account struct {
	AccountNumber int64
	Name          string
	Balance       float64
	Transactions  []Transaction
}

type Transaction struct {
	TransactionID int32
	Type          string
	Amount        float64
	Timestamp     time.Time
}

// map to track and update the
var accounts = make(map[int64]*Account) 

//// Account created successfully: Name: Oyindamola Abiola, Account number: 1501629802, Account balance: 400000.00 

//// Account created successfully: Name: Efunroye Dasola, Account number: 1041001538, Account balance: 698548.00 - 2900

// all account
var allAccount = []Account{}

// transaction slice
var transactionList = []Transaction{}

/*  the value(*Account) of the map is a pointer, pointing to the newAccount
is the memory storage of every newly created account, as a reference.
*/

func createAccount(accountName string, initialDeposit float64) (*Account, error) {

	// checks for name == ""
	if accountName == "" {
		return nil, errors.New("name cannot be empty")
	}

	// checks for the initial deposit != (-neg) || 0
	if initialDeposit <= 0 {
		return nil, fmt.Errorf("the amount must be greater than zero: %.2f", initialDeposit)
	}

	// generate the random account number
	seed := rand.NewSource(time.Now().UTC().UnixNano())
	source := rand.New(seed)
	accountNumber := source.Int63n(999999999) + 1000000000

	// check if account exists

	// create new account number
	newAccount := &Account{
		AccountNumber: accountNumber,
		Name:          accountName,
		Balance:       initialDeposit,
		Transactions:  []Transaction{},
	}

	// save in maps pointing to the memory location of the new accounts
	accounts[accountNumber] = newAccount

	allAccount = append(allAccount, *newAccount)

	// set the transaction struct
	initialTxn := Transaction{
		TransactionID: int32(len(newAccount.Transactions) + 1),
		Type:          "Deposit",
		Amount:        initialDeposit,
		Timestamp:     time.Now(),
	}

	// update the user record transactions
	newAccount.Transactions = append(newAccount.Transactions, initialTxn)

	// add transaction globally
	transactionList = append(transactionList, initialTxn)

	fmt.Printf("Account created successfully: Name: %s, Account number: %d, Account balance: %f \n",
		newAccount.Name, newAccount.AccountNumber, newAccount.Balance)

	return newAccount, nil
}

func depositMoney(accountNumber int64, amount float64) (*Account, error) {
	// check for the account existence
	if _, exists := accounts[accountNumber]; !exists {
		return nil, fmt.Errorf("account number %d doesn't exist", accountNumber)
	}

	// checks for the amount
	if amount <= 0 {
		return nil, fmt.Errorf("the amount %.2f you entered must be greater than zero", amount)
	}

	account := accounts[accountNumber]
	account.Balance += amount

	// set the transaction struct
	depositTxn := Transaction{
		TransactionID: int32(len(account.Transactions) + 1),
		Type:          "Deposit",
		Amount:        amount,
		Timestamp:     time.Now(),
	}

	// update the user record transactions
	account.Transactions = append(account.Transactions, depositTxn)

	// add transaction globally
	transactionList = append(transactionList, depositTxn)

	fmt.Printf("Deposit of %.2f into account %d is successful \n", amount, accountNumber)

	return account, nil
}

func withdrawMoney(accountNumber int64, amount float64) error {
	// check for the account existence
	if _, exists := accounts[accountNumber]; !exists {
		return fmt.Errorf("Account number %d doesn't exist", accountNumber)
	}

	// checks for the amount
	if amount <= 0 {
		return fmt.Errorf("the amount %.2f you entered must be greater than zero", amount)
	}

	account := accounts[accountNumber]
	account.Balance -= amount

	// set the transaction struct
	withdrawTxn := Transaction{
		TransactionID: int32(len(account.Transactions) + 1),
		Type:          "Withdraw",
		Amount:        amount,
		Timestamp:     time.Now(),
	}

	// record the transaction for the user
	account.Transactions = append(account.Transactions, withdrawTxn)

	// record the transaction global
	transactionList = append(transactionList, withdrawTxn)

	fmt.Printf("Withdrawal of %.2f from account %d is successful. \n", amount, accountNumber)

	return nil
}

func transferMoney(sender int64, receiver int64, amount float64) error {
	// check for the accounts existence
	if _, exists := accounts[sender]; !exists {
		return fmt.Errorf("account number %d doesn't exist", sender)
	}

	if _, exists := accounts[receiver]; !exists {
		return fmt.Errorf("account number %d doesn't exist", receiver)
	}

	// checks for the amount
	if amount <= 0 {
		return fmt.Errorf("the amount %.2f you entered must be greater than zero", amount)
	}

	senderAccount := accounts[sender]
	receiverAccount := accounts[receiver]

	// set transaction for the sender
	senderTxn := Transaction{
		TransactionID: int32(len(senderAccount.Transactions) + 1),
		Type:          "Transfer",
		Amount:        amount,
		Timestamp:     time.Now(),
	}

	// record this transaction for the user
	senderAccount.Transactions = append(senderAccount.Transactions, senderTxn)

	// record this transaction
	transactionList = append(transactionList, senderTxn)

	senderAccount.Balance -= amount
	receiverAccount.Balance += amount

	// set the transaction struct
	receiverTxn := Transaction{
		TransactionID: int32(len(receiverAccount.Transactions) + 1),
		Type:          "Credit",
		Amount:        amount,
		Timestamp:     time.Now(),
	}

	// set the transaction for the user
	receiverAccount.Transactions = append(receiverAccount.Transactions, receiverTxn)

	// record globally
	transactionList = append(transactionList, receiverTxn)

	fmt.Printf("Transfer of %.2f from account %d to account %d is successful \n", amount, sender, receiver)

	return nil
}

func viewAccountDetails(accountNumber int64) (*Account, error) {
	account, exists := accounts[accountNumber]
	if !exists {
		return nil, fmt.Errorf("Account number %d is either invalid or doesn't exist", accountNumber)

	}

	// prints the struct fields with their names
	fmt.Printf("Account Details: %+v \n", account)

	return account, nil
}

func generateStatement(accountNumber int64, ) error {
	account, exists := accounts[accountNumber]
	if !exists {
		return fmt.Errorf("Account number %d is either invalid or doesn't exist", accountNumber)

	}

	accountStatement := account.Transactions

	fmt.Printf("Statement of the account: Name: %s, %+v \n", account.Name, accountStatement)

	return nil
}

func filter(accountNumber int64, fromDate, toDate time.Time) ([]Transaction, error) {
	account, exists := accounts[accountNumber]
	if !exists {
		return nil, fmt.Errorf("Account number %d is either invalid or doesn't exist", accountNumber)
	}

	// Filter transactions based on the date range
	var filteredTransactions []Transaction
	for _, txn := range account.Transactions {
		if txn.Timestamp.After(fromDate) && txn.Timestamp.Before(toDate) || txn.Timestamp.Equal(fromDate) || txn.Timestamp.Equal(toDate) {
			filteredTransactions = append(filteredTransactions, txn)
		}
	}

	if len(filteredTransactions) == 0 {
		return nil, fmt.Errorf("No transactions found for Account %d within the specified range", accountNumber)
	}

	return filteredTransactions, nil
}

func displayAllAccounts() {
	for _, account := range accounts {
		fmt.Printf("Accounts: %+v \n", account)
	}
}

func main() {
	// CREATE ACCOUNT
	// getting different account number for the same user at every call ????
	account, err := createAccount("Oyindamola Abiola", 1000000.00)
	if err != nil { // Check for error
		fmt.Println("Error:", err)
		return
	}

	// account 2
	account2, err2 := createAccount("Efunroye Abosede", 200000.00)
	if err2 != nil { // Check for error
		fmt.Println("Error:", err2)
		return
	}

	// DEPOSIT
	_, depositErr := depositMoney(account.AccountNumber, 700000000.00)
	if depositErr != nil {
		fmt.Println("Error:", depositErr)
		return
	}

	// WITHDRAW
	amountWithdrawn := 20000.00
	withdrawErr := withdrawMoney(account.AccountNumber, amountWithdrawn)
	if withdrawErr != nil {
		fmt.Println("Error:", withdrawErr)
		return
	}

	// TRANSFER
	amountTransferred := 90000.00
	transferErr := transferMoney(account.AccountNumber, account2.AccountNumber, amountTransferred)
	if transferErr != nil { // not empty
		fmt.Println("Error: ", transferErr)
		return
	}

	// ACCOUNT DETAILS
	account, accDetailErr := viewAccountDetails(account2.AccountNumber)
	if accDetailErr != nil {
		fmt.Println("Error: ", accDetailErr)
		return
	}

	// STATEMENT
	accountStatementErr := generateStatement(account.AccountNumber)
	if accountStatementErr != nil {
		fmt.Println("Error: ", accountStatementErr)
		return
	}

	// ACCOUNTS DISPLAY
	displayAllAccounts()
}



// 				Simple Banking System Objective
// Build a simple banking system using Go that allows users to manage accounts, and perform transactions. This project will help you solidify your understanding of Go's core concepts, including structs, interfaces and error handling.
// Requirements
// 1. Functional Requirements
//     Account Management:
//     Create a new account.
//     View account details (Account Number, Name, Balance).
// Update account holder's details.
//     Transactions:
//         Deposit money into an account.
//         Withdraw money from an account.
//         Transfer money between accounts.
// Error Handling:
//     Handle scenarios like insufficient balance, invalid account number, or incorrect input.
//     User-Friendly Interface:
//     Provide a text-based menu-driven interface for interaction.
// Features to Implement
// 1. Core Structs
//     Account:
//         Fields: AccountNumber, Name, Balance, Transactions.
//         Methods: Deposit, Withdraw, Transfer.
//     Transaction:
//         Fields: TransactionID, Type (Deposit/Withdrawal/Transfer), Amount, Timestamp.
// 2. Operations
//     Create Account:
//     Generate a unique account number.
//     Accept user details and initial deposit.
// Deposit:
//     Increase the account balance.
//     Record the transaction.
// Withdraw:
//     Decrease the account balance if sufficient funds are available.
//     Record the transaction.
// Transfer:
//     Deduct the amount from the sender's account.
//     Add the amount to the receiver's account.
//     Record the transaction in both accounts.
// View Account Details:
//     Show account information and balance.
// Display All Accounts:
//     Show a summary of all accounts and their balances.
// Project Phases
// Phase 1: Setup and Account Management
//     Create a main.go file with a basic menu-driven interface.
//     Define structs for Account and Transaction.
//     Implement functions to:
//     Create accounts.
//     View account details.
// Phase 2: Transactions
//     Implement Deposit and Withdraw methods for Account.
//     Add transaction records for each operation.
// Phase 6: Error Handling
//     Handle invalid inputs and operations (e.g., withdrawal from a non-existent account).
//     Implement proper error messages and logging.
// Deliverables
//     A fully functional program with:
//     Menu-driven interface.
//     Well-commented code.




// Here is a starter code to help you jumpstart (edited) 
// package main
// import (
//     "fmt"
//     "time"
// )
// func main() {
//     type Account struct {
//         AccountNumber int
//         Name          string
//         Balance       float64
//         Transactions  []Transaction
//     }
    
//     type Transaction struct {
//         TransactionID string
//         Type          string
//         Amount        float64
//         Timestamp     time.Time
//     }
    
//     for {
//         fmt.Println("\n=== Banking System ===")
//         fmt.Println("1. Create Account")
//         fmt.Println("2. Deposit Money")
//         fmt.Println("3. Withdraw Money")
//         fmt.Println("4. Transfer Money")
//         fmt.Println("5. View Account Details")
//         fmt.Println("6. Generate Account Statement")
//         fmt.Println("7. Display All Accounts")
//         fmt.Println("8. Exit")
//         var choice int
//         fmt.Print("Enter your choice: ")
//         fmt.Scan(&choice)
//         switch choice {
//         case 1:
//             createAccount()
//         case 2:
//             depositMoney()
//         case 3:
//             withdrawMoney()
//         case 4:
//             transferMoney()
//         case 5:
//             viewAccountDetails()
//         case 6:
//             generateStatement()
//         case 7:
//             displayAllAccounts()
//         case 8:
//             fmt.Println("Exiting... Thank you!")
//             return
//         default:
//             fmt.Println("Invalid choice. Please try again.")
//         }
//     }
// }

// func createAccount() {
//     fmt.Println("Creating account...")
//     // Implement functionality
// }
// func depositMoney() {
//     fmt.Println("Depositing money...")
//     // Implement functionality
// }
// func withdrawMoney() {
//     fmt.Println("Withdrawing money...")
//     // Implement functionality
// }
// func transferMoney() {
//     fmt.Println("Transferring money...")
//     // Implement functionality
// }
// func viewAccountDetails() {
//     fmt.Println("Viewing account details...")
//     // Implement functionality
// }
// func generateStatement() {
//     fmt.Println("Generating account statement...")
//     // Implement functionality
// }
// func displayAllAccounts() {
//     fmt.Println("Displaying all accounts...")
//     // Implement functionality
// }


// January 9TH !!!
// All accounts created should be stored in a JSON file, accessing, deposit, updates...
// every function call should also be called from the Json file
// When generating bank statement it should be in an excel sheet// 