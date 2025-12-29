# Interest Calculator

A Go program for calculating interest rates and projecting balance growth over time.

## Features

- Calculate interest rates based on account balance
- Compute interest amounts and annual balance updates
- Determine years needed to reach a target balance

## Interest Rate Tiers

| Balance Range | Interest Rate |
|---------------|---------------|
| $0 - $999.99 | 0.5% |
| $1,000 - $4,999.99 | 1.621% |
| $5,000+ | 2.475% |
| Negative balances | 3.213% |

## Functions

### `InterestRate(balance float64) float32`
Returns the applicable interest rate for a given balance.

### `Interest(balance float64) float64`
Calculates the interest amount for a balance.

### `AnnualBalanceUpdate(balance float64) float64`
Calculates the updated balance after one year with interest.

### `YearsBeforeDesiredBalance(balance, targetBalance float64) int`
Determines the minimum years required to reach a target balance.

## Usage Example

```go
balance := 200.75
rate := InterestRate(balance)    // 0.5
interest := Interest(balance)    // 1.00375
newBalance := AnnualBalanceUpdate(balance)  // 201.75375

years := YearsBeforeDesiredBalance(200.75, 214.88)  // 14 years
```

## Note

This implementation uses compound interest calculation, where interest is added to the balance annually for subsequent years.