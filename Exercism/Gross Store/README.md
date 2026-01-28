# Gross Store Inventory Management

A Go package for managing inventory and billing in a gross store using traditional bulk measurement units.

## Description

This package provides functionality to manage customer bills and inventory using traditional bulk measurement units like dozen, gross, and great gross. It supports adding items, removing items, and checking quantities in a customer's bill.

## Functions

### `Units() map[string]int`
Returns a map of available bulk measurement units with their corresponding quantities.

**Returns:**
- `map[string]int`: A map containing unit names as keys and their quantities as values.

**Available units:**
- `"quarter_of_a_dozen"`: 3 items
- `"half_of_a_dozen"`: 6 items
- `"dozen"`: 12 items
- `"small_gross"`: 120 items
- `"gross"`: 144 items
- `"great_gross"`: 1728 items

### `NewBill() map[string]int`
Creates a new empty customer bill.

**Returns:**
- `map[string]int`: An empty map representing a new bill.

### `AddItem(bill, units map[string]int, item, unit string) bool`
Adds an item to the customer's bill in the specified unit quantity.

**Parameters:**
- `bill` (map[string]int): The customer's current bill
- `units` (map[string]int): The map of available units (from `Units()`)
- `item` (string): The name of the item to add
- `unit` (string): The unit of measurement for the item

**Returns:**
- `bool`: `true` if the item was successfully added, `false` if the unit is invalid.

**Example:**
```go
AddItem(bill, units, "carrot", "dozen") // Adds 12 carrots to the bill
```

### `RemoveItem(bill, units map[string]int, item, unit string) bool`
Removes an item from the customer's bill in the specified unit quantity.

**Parameters:**
- `bill` (map[string]int): The customer's current bill
- `units` (map[string]int): The map of available units (from `Units()`)
- `item` (string): The name of the item to remove
- `unit` (string): The unit of measurement for the item

**Returns:**
- `bool`: `true` if the item was successfully removed, `false` if:
    - The item doesn't exist in the bill
    - The unit is invalid
    - Removing would result in negative quantity

**Behavior:**
- If the new quantity becomes 0, the item is completely removed from the bill.
- If the new quantity is positive, the bill is updated with the reduced quantity.

### `GetItem(bill map[string]int, item string) (int, bool)`
Retrieves the quantity of a specific item from the customer's bill.

**Parameters:**
- `bill` (map[string]int): The customer's current bill
- `item` (string): The name of the item to check

**Returns:**
- `int`: The quantity of the item (0 if item doesn't exist)
- `bool`: `true` if the item exists in the bill, `false` otherwise

## Usage Examples

```go
package main

import (
    "fmt"
)

func main() {
    // Get available units
    units := Units()
    fmt.Println("Available units:", units)
    
    // Create a new bill
    bill := NewBill()
    fmt.Println("New bill:", bill)
    
    // Add items to the bill
    success := AddItem(bill, units, "carrot", "dozen")
    fmt.Println("Added carrots (dozen):", success) // true
    
    success = AddItem(bill, units, "apple", "half_of_a_dozen")
    fmt.Println("Added apples (half dozen):", success) // true
    
    success = AddItem(bill, units, "banana", "invalid_unit")
    fmt.Println("Added with invalid unit:", success) // false
    
    // Check current quantities
    qty, exists := GetItem(bill, "carrot")
    fmt.Printf("Carrots: %d (exists: %v)\n", qty, exists) // Carrots: 12 (exists: true)
    
    // Remove items from the bill
    success = RemoveItem(bill, units, "carrot", "quarter_of_a_dozen")
    fmt.Println("Removed 3 carrots:", success) // true
    
    qty, exists = GetItem(bill, "carrot")
    fmt.Printf("Carrots after removal: %d\n", qty) // Carrots after removal: 9
    
    // Try to remove more than available
    success = RemoveItem(bill, units, "carrot", "dozen")
    fmt.Println("Tried to remove 12 carrots (only 9 left):", success) // false
    
    // Remove all remaining carrots
    success = RemoveItem(bill, units, "carrot", "half_of_a_dozen")
    fmt.Println("Removed 6 carrots:", success) // true
    
    success = RemoveItem(bill, units, "carrot", "quarter_of_a_dozen")
    fmt.Println("Removed final 3 carrots:", success) // true
    
    qty, exists = GetItem(bill, "carrot")
    fmt.Printf("Carrots after complete removal: %d (exists: %v)\n", qty, exists) // 0, false
}
```

## Unit Conversion Reference

| Unit | Quantity | Equivalent To |
|------|----------|--------------|
| quarter_of_a_dozen | 3 | ¼ dozen |
| half_of_a_dozen | 6 | ½ dozen |
| dozen | 12 | 12 items |
| small_gross | 120 | 10 dozen |
| gross | 144 | 12 dozen |
| great_gross | 1728 | 12 gross |

## Error Handling

All functions that can fail return a boolean value indicating success:
- `AddItem()` returns `false` for invalid units
- `RemoveItem()` returns `false` for:
    - Non-existent items
    - Invalid units
    - Attempts to remove more than available quantity
- `GetItem()` returns `(0, false)` for non-existent items

## Sample Output

```
Available units: map[quarter_of_a_dozen:3 half_of_a_dozen:6 dozen:12 small_gross:120 gross:144 great_gross:1728]
New bill: map[]
Added carrots (dozen): true
Added apples (half dozen): true
Added with invalid unit: false
Carrots: 12 (exists: true)
Removed 3 carrots: true
Carrots after removal: 9
Tried to remove 12 carrots (only 9 left): false
Removed 6 carrots: true
Removed final 3 carrots: true
Carrots after complete removal: 0 (exists: false)
```

## Dependencies

- Go 1.13 or higher
- No external packages required

## Use Cases

- Retail inventory management
- Bulk goods tracking
- Wholesale order processing
- Grocery store billing systems