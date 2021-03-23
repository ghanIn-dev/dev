package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Title string
	Body  string
}

var db *gorm.DB
var err error

func (p *Page) save() {
	result := db.First(&Page{}, "title = ?", p.Title)
	if result.RowsAffected != 0 {
		db.Model(&Page{}).Where("title = ?", p.Title).Update("body", p.Body)
	} else {
		db.Create(&Page{Title: p.Title, Body: p.Body})
	}
}

func loadPage(title string) (*Page, error) {
	var p Page
	result := db.First(&p, "title = ?", title)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &Page{Title: p.Title, Body: p.Body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	// 테이블 자동 생성
	db.AutoMigrate(&Page{})

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
