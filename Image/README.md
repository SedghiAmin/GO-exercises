# Custom Image Generator in Go

A Go program that generates custom images with various patterns by implementing the `image.Image` interface.

## Features

- Generates images with different patterns:
  - **Stripes**: Vertical stripes pattern
  - **Gradient**: Diagonal gradient pattern
  - **Circle**: Circular pattern based on distance from center
  - **Default**: Checkerboard-like pattern

- Saves images as PNG files
- Customizable dimensions and patterns
- Implements the standard `image.Image` interface

## Pattern Types

1. **`stripes`**: Creates vertical stripes where color intensity varies with X coordinate
2. **`gradient`**: Creates a diagonal gradient where color intensity varies with (X+Y)/2
3. **`circle`**: Creates circular patterns based on distance from origin (0,0)
4. **default**: Creates a pattern where color intensity varies with X*Y

## Color Scheme

All patterns use a blue-based color scheme:
- Red and Green components: Determined by the pattern algorithm
- Blue component: Always maximum (255)
- Alpha component: Always maximum (255) - fully opaque

## Usage

### Basic Usage
```go
// Create a 256x256 image with circle pattern
m := Image{256, 256, "circle"}

Changing Patterns
// Vertical stripes
m := Image{256, 256, "stripes"}

// Diagonal gradient  
m := Image{256, 256, "gradient"}

// Circular pattern
m := Image{256, 256, "circle"}

Custom Dimensions
// Create a 512x512 image
m := Image{512, 512, "gradient"}

// Create a 128x128 image
m := Image{128, 128, "stripes"}
```
# Requirements
Go 1.16 or higher

Standard library only (no external dependencies)

# Installation
Clone or download the source code

Navigate to the directory containing main.go

# Run the program:

bash
go run main.go

# Output
The program generates a PNG file named my_image.png in the same directory with the specified pattern and dimensions.

# How It Works
The program implements the image.Image interface by defining three required methods:

ColorModel(): Returns the color model (RGBA)

Bounds(): Returns the image dimensions

At(x, y): Calculates the color for each pixel based on the selected pattern

# Extending
You can easily add new patterns by:

Adding a new case in the At() method

Implementing your own color calculation algorithm

Adding new configuration options to the Image struct

# Example Output
Circle pattern: Creates concentric blue circles

Stripes pattern: Creates vertical blue stripes of varying intensity

Gradient pattern: Creates a smooth diagonal blue gradient

# License
This code is provided as-is for educational purposes.