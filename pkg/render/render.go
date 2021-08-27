package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap {

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w, tmpl)
	if err != nil {fmt.Println("Error getting template cache.")}
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.gohtml")
	//Execute sending template to web browser
	err = parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func RenderTemplateTest(w http.ResponseWriter, tmpl string) (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)

	//Create a map with index<->template
	myCache := map[string]*template.Template{}

	//get the gohtml that match the pattern
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {return myCache, err}

	for _, page := range pages {
		//get the name of the template page
		name := filepath.Base(page)

		log.Println("Rendering page", name)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {return myCache, err}

		//find any layout match the template
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {return myCache, err}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {return myCache, err}
		}
		myCache[name] = ts
	}
	return myCache, err
}
