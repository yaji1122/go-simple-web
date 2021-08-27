package handler

import (
	"github.com/yaji1122/go-simple-web/pkg/render"
	"net/http"
)

//Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}
//func Home(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "This is the home page.")
//}

//About page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
//func About(w http.ResponseWriter, r *http.Request) {
//	sum, _ := addValues(2, 3)
//	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 3 is %d", sum))
//}


