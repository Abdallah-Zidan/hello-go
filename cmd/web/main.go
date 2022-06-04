package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Abdallah-Zidan/hello-go/pkg/config"
	"github.com/Abdallah-Zidan/hello-go/pkg/handlers"
	"github.com/Abdallah-Zidan/hello-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var session *scs.SessionManager = scs.New()

var app = config.AppConfig{
	TemplateDir:  "./templates",
	InProduction: false,
	UseCache:     false,
	Session:      session,
}

func main() {
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Path = "/"
	session.Cookie.Secure = app.InProduction
	session.Cookie.HttpOnly = true
	session.Cookie.SameSite = http.SameSiteStrictMode

	tc, err := render.CacheTemplates("templates")

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	repo := handlers.NewRepository(&app)

	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Println("Server Listening on port " + port)
	log.Fatal(srv.ListenAndServe())
}
