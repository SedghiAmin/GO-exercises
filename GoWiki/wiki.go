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

func loadPage(title string) (*Page, error){
	fileName:= title + ".txt"
	body, err := os.ReadFile(fileName) //The standard library function os.ReadFile returns []byte and error.
	if err != nil{
		return nil, err
	}
	return &Page{title, body}, nil
}

func main(){
	p1:= &Page{Title: "test", Body: []byte("This is a test file.")}
	p1.Save()

	p2, _:= loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}