package main

import "fmt"

type BankAccount struct {
    owner   string
    balance float64
}

// ✅ METHOD - Account related operations
func (a *BankAccount) Deposit(amount float64) {
    a.balance += amount
    fmt.Printf("%s deposited %.2f\n", a.owner, amount)
}

func (a *BankAccount) Withdraw(amount float64) bool {
    if a.balance >= amount {
        a.balance -= amount
        fmt.Printf("%s withdrew %.2f\n", a.owner, amount)
        return true
    }
    fmt.Println("Insufficient funds!")
    return false
}

func (a BankAccount) GetBalance() float64 {
    return a.balance
}

// ✅ FUNCTION - Operation between two accounts
func Transfer(from, to *BankAccount, amount float64) bool {
    if from.Withdraw(amount) {
        to.Deposit(amount)
        return true
    }
    return false
}

func main() {
    acc1 := &BankAccount{owner: "Amin", balance: 1000}
    acc2 := &BankAccount{owner: "Ali", balance: 500}
    
    // METHOD calls
    acc1.Deposit(200)
    fmt.Printf("Amin's balance: %.2f\n", acc1.GetBalance())
    
    // FUNCTION call
    Transfer(acc1, acc2, 300)
    
    fmt.Printf("Amin's balance: %.2f\n", acc1.GetBalance())
    fmt.Printf("Ali's balance: %.2f\n", acc2.GetBalance())
}
