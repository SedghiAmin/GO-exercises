# Bank Account Example in Go

This repository demonstrates how to implement a simple **BankAccount** system in Go using both **methods** and **functions**. The example shows how pointer receivers allow methods to modify the original struct, how value receivers are used for read-only operations, and how standalone functions can coordinate actions between multiple accounts.

## Overview

The program defines a `BankAccount` struct with the following fields:

* `owner`: Name of the account holder
* `balance`: Current account balance

It also implements:

* **Methods** for depositing, withdrawing, and retrieving the balance
* A **function** for transferring money between two accounts

This demonstrates Go's method receivers and when to use pointer versus value receivers.

## Code

```go
package main

import "fmt"

type BankAccount struct {
    owner   string
    balance float64
}

// ✅ METHOD - Account-related operations
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
```

## Output

```
Amin deposited 200.00
Amin's balance: 1200.00
Amin withdrew 300.00
Ali deposited 300.00
Amin's balance: 900.00
Ali's balance: 800.00
```

## Key Concepts

### ✅ Pointer Receivers

Used for `Deposit` and `Withdraw` because these operations modify the account balance.

### ✅ Value Receiver

Used for `GetBalance` because it only reads the balance and does not modify the struct.

### ✅ Function vs Method

* **Methods** are tied to a single account instance.
* **Functions** like `Transfer` operate **between multiple accounts**, making them independent of any one instance.

## Summary

This example illustrates how Go handles struct methods, when to use pointers, and how to structure logic that involves multiple objects. The combination of methods and standalone functions leads to clean, understandable, and maintainable Go code.
