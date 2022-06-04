package render

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Abdallah-Zidan/hello-go/pkg/config"
	"github.com/Abdallah-Zidan/hello-go/pkg/models"
)

var app *config.AppConfig

func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

var functions = template.FuncMap{}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	if td == nil {
		td = &models.TemplateData{
			StringMap: make(map[string]string),
		}
	} else if td.StringMap == nil {
		td.StringMap = make(map[string]string)
	}

	_, ok := td.StringMap["pageTitle"]

	if !ok {
		td.StringMap["pageTitle"] = "Page Title"
	}

	return td
}

func RenderTemplate(file string, w http.ResponseWriter, data *models.TemplateData) {

	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CacheTemplates(app.TemplateDir)
		if err != nil {
			onError(w, err)
		}
	}

	tmpl, ok := tc[file]

	if !ok {
		onError(w, errors.New("template not found in cache"))
	}

	data = AddDefaultData(data)

	err = tmpl.Execute(w, data)

	if err != nil {
		onError(w, err)
	}
}

func CacheTemplates(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(dir + "/*.page.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob(dir + "/*.layout.html")

		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseFiles(matches...)
			if err != nil {
				return nil, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}

func onError(w http.ResponseWriter, err error) {
	log.Println("faild to render template", err.Error())
	w.WriteHeader(500)
	w.Write([]byte("Internal Server Error"))
}
