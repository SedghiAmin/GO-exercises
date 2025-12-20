# Card Tricks

This exercise focuses on working with slices in Go.  
You will practice common slice operations such as accessing elements, updating values, appending, prepending, and removing items.

## Functions

### `FavoriteCards() []int`
Returns a slice containing your favorite cards in a fixed order: `2, 6, 9`.

### `GetItem(slice []int, index int) int`
Retrieves the item at the given index from the slice.
- Returns `-1` if the index is out of range.

### `SetItem(slice []int, index, value int) []int`
Updates the item at the given index with a new value.
- If the index is out of range, the value is appended to the slice.

### `PrependItems(slice []int, values ...int) []int`
Adds one or more values to the beginning of the slice.

### `RemoveItem(slice []int, index int) []int`
Removes the item at the specified index.
- If the index is out of range, the original slice is returned unchanged.

## Example

```go
FavoriteCards()                   // []int{2, 6, 9}
GetItem([]int{1, 2, 4, 1}, 10)     // -1
SetItem([]int{1, 2, 4, 1}, -1, 6)  // []int{1, 2, 4, 1, 6}
PrependItems([]int{3, 2, 6}, 5, 1) // []int{5, 1, 3, 2, 6}
RemoveItem([]int{3, 2, 6, 4}, 2)   // []int{3, 2, 4}
```
# Learning Goals

- Understanding slice indexing and bounds checking

- Modifying slices safely

- Using append with variadic arguments

- Applying idiomatic Go slice patterns