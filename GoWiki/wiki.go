package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"errors"
)

//The function template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the *Template unaltered. 
var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte // that is the type expected by the io libraries we will use
}

func (p *Page) Save() error { // returns an error value because that is the return type of WriteFile
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) //The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be created with read-write permissions for the current user only.
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error){
	m:= validPath.FindStringSubmatch(r.URL.Path)
	if m == nil{
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName) //The standard library function os.ReadFile returns []byte and error.

	if err != nil {
		return nil, err
	}

	return &Page{title, body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err:= getTitle(w, r)

	if err != nil{
		return
	}

	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err:= getTitle(w, r)
	
	if err != nil{
		return
	}

	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err:= getTitle(w, r)
	
	if err != nil{
		return
	}
	
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	err = p.Save()

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
}

func main() {
	/* p1 := &Page{Title: "test", Body: []byte("This is a test file.")}
	p1.Save() */

	/* p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body)) */

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) //With this web server running, a visit to http://localhost:8080/view/test should show a page titled "test" containing the words "Hello world".

}
