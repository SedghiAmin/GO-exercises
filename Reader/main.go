package main

import(
	"fmt"
	"strings"
	"io"
)

func main() {
	str:= "This is a sample text. please read me ...!";
	source := strings.NewReader(str)
	buffer := make([]byte, 8)
	
	iteration := 1

	for{
		n, err := source.Read(buffer)

		fmt.Printf("=== iteration: %d ===\n", iteration)
		fmt.Printf("Number of bytes read: %d\n", n)
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Buffer contents: %v\n", buffer)
		fmt.Printf("Valid data: %q\n", buffer[:n])
		fmt.Printf("Valid data (string): %s\n", string(buffer[:n]))
		fmt.Println()
		
		if err == io.EOF {
			fmt.Println("We have reached the end of the data.")
			break
		}

		iteration += 1
	}

}