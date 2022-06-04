package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLogger    *log.Logger
	TemplateDir   string
	InProduction  bool
	Session       *scs.SessionManager
}
