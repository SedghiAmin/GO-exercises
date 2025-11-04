package main

// External package from Go Tour used for displaying generated images
// Install with: go get golang.org/x/tour/pic
import "golang.org/x/tour/pic"

// Function that returns a 2D slice of uint8 values
// This function is called by pic.Show
func Pic(dx, dy int) [][]uint8 {
	// Create the main slice with dy rows
	picture := make([][]uint8, dy)
	
	// For each row
	for y := 0; y < dy; y++ {
		// Create a slice with dx columns
		picture[y] = make([]uint8, dx)
		
		// Fill each pixel in this row
		for x := 0; x < dx; x++ {
			// Different formulas for generating patterns
			
			// Option 1: Average of x and y (simple gradient)
			// picture[y][x] = uint8((x + y) / 2)
			
			// Option 2: Multiply x and y (interesting pattern)
			picture[y][x] = uint8(x ^ y)
			
			// Option 3: XOR between x and y (geometric pattern)
			// picture[y][x] = uint8(x ^ y)
			
			// Option 4: Combination (very cool)
			// picture[y][x] = uint8((x * y) / 2)
			
			// Option 5: Wave pattern
			// picture[y][x] = uint8(x % 256)
			
			// Option 6: Checkerboard pattern
			// picture[y][x] = uint8((x/10 + y/10) % 2 * 255)
		}
	}
	
	return picture
}

func main() {
	pic.Show(Pic)
}
