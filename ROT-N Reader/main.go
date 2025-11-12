package main

import(
	"strings"
	"io"
	"os"
)

type rotReader struct{
	r io.Reader
	n byte
}

func (rR rotReader) Read(b []byte) (int , error){
	// Step 1: Read data from the underlying reader
	n, err := rR.r.Read(b)

	if err != nil{
		return n, err
	}

	// Step 2: Apply ROT to each character
	for i := range n {
		p:= b[i]

		// Handle uppercase letters
		if p >= 'A' && p <= 'Z'{
			// Shift by n positions, wrap around using modulo 26
			b[i] = (p - 'A' + rR.n) % 26 + 'A'
		}
		// Handle lowercase letters
		if p >= 'a' && p <= 'z'{
			b[i] = (p - 'a' + rR.n) % 26 + 'a'
		}
		// Non-letters remain unchanged
	}

	return n ,nil

}

func main(){
	// The encrypted message
	src := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rotReader{src, 13}
	
	// This will print: "You cracked the code!"
	io.Copy(os.Stdout, &r)

	// Let's test with more examples
	println("\n\n=== Additional Tests ===")

	testCases := []string{
		"Hello World!" ,
		"ROT13 example" ,
		"abcxyz ABCXYZ" ,
	}

	for _, txt:= range testCases{
		r:= rotReader{strings.NewReader(txt), 1}
		io.Copy(os.Stdout, &r)
		println()
	}
}