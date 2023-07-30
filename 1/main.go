package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/caarlos0/env"
)

type App struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func (a *App) ParseConfig() error {
	return env.Parse(a)
}

func (a *App) Run() error {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("healthy", r.URL.Path, r.Method, r.RemoteAddr, r.UserAgent())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})
	return http.ListenAndServe(":"+a.Port, serverMux)
}

func NewApp() *App {
	app := &App{}
	return app
}

func main() {
	app := NewApp()
	err := app.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on port %s", app.Port)
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
