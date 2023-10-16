package render

import (
	"html/template"
	"log"
	"net/http"
)

var TemplateCache map[string]*template.Template

type tempData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	BoolMap   map[string]bool
	CustomMap map[string]interface{}
}

func RenderTemplate(w http.ResponseWriter, tmpName string, data any) error {
	tpl, ok := TemplateCache[tmpName]
	if !ok {
		log.Println("no template found in cache")
		tpl, err := template.ParseFiles("./templates/" + tmpName + ".html")
		if err != nil {
			return err
		}
		TemplateCache[tmpName] = tpl
		err = tpl.Execute(w, nil)
		if err != nil {
			return err
		}

	} else {
		log.Println("using cached templated")
		err := tpl.Execute(w, data)
		if err != nil {
			return err
		}
	}
	return nil
}
