package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math"
	"fmt"
)

// Image represents a custom image with configurable patterns
type Image struct {
	width, height int
	patternType   string
}

// Bounds returns the dimensions of the image
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel returns the color model (RGBA)
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the color of the pixel at position (x, y)
func (img Image) At(x, y int) color.Color {
	var v uint8
	
	switch img.patternType {
	case "stripes":
		v = uint8(x) // Vertical stripes
	case "gradient":
		v = uint8((x + y) / 2) // Diagonal gradient
	case "circle":
		v = uint8(math.Sqrt(float64(x*x + y*y))) // Circular pattern
	default:
		v = uint8(x * y) // Default pattern
	}
	
	return color.RGBA{v, v, 255, 255} // Blue-based colors
}

func main() {
	// Create a 256x256 image with circle pattern
	m := Image{256, 256, "circle"}
	
	// Create output file
	file, err := os.Create("my_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// Encode and save as PNG
	err = png.Encode(file, m)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("✅ Image successfully saved to 'my_image.png'")
}