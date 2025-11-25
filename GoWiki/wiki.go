package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte // that is the type expected by the io libraries we will use
}

func (p *Page) Save() error { // returns an error value because that is the return type of WriteFile
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) //The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be created with read-write permissions for the current user only.
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
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	err:= p.Save()

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	err = t.Execute(w, p)
	
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
