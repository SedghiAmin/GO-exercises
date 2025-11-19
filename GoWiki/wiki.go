package main

import (
	"fmt"
	"os"
)

type Page struct{
	Title string
	Body []byte // that is the type expected by the io libraries we will use
}

func (p *Page) Save() error{	// returns an error value because that is the return type of WriteFile
	filename:= p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) //The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be created with read-write permissions for the current user only.
}

