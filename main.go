package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	TEMPLATES_RECOVERY     string `env:"GOTRUE_MAILER_TEMPLATES_RECOVERY" envDefault:"./templates/recovery.html"`
	TEMPLATES_INVITE       string `env:"GOTRUE_MAILER_TEMPLATES_INVITE" envDefault:"./templates/invite.html"`
	TEMPLATES_CONFIRMATION string `env:"GOTRUE_MAILER_TEMPLATES_CONFIRMATION" envDefault:"./templates/confirmation.html"`
	TEMPLATES_MAGIC_LINK   string `env:"GOTRUE_MAILER_TEMPLATES_MAGIC_LINK" envDefault:"./templates/magic-link.html"`
	TEMPLATES_EMAIL_CHANGE string `env:"GOTRUE_MAILER_TEMPLATES_EMAIL_CHANGE" envDefault:"./templates/email-change.html"`
}

func main() {
	godotenv.Load()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("Gotrue email template server starting...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the gotrue email template server")
	})

	http.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.TEMPLATES_RECOVERY)
	})

	http.HandleFunc("/invite", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.TEMPLATES_INVITE)
	})

	http.HandleFunc("/confirmation", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.TEMPLATES_CONFIRMATION)
	})

	http.HandleFunc("/magic-link", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.TEMPLATES_MAGIC_LINK)
	})

	http.HandleFunc("/email-change", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.TEMPLATES_EMAIL_CHANGE)
	})

	http.ListenAndServe(":3000", nil)
}
