package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.com/deliveos/i18n"
	"gitlab.com/deliveos/webappbot/config"
	"gitlab.com/deliveos/webappbot/internal/webappbot"
	"gopkg.in/yaml.v3"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Localization
	localizer := i18n.NewLocalizer("en", "internal/i18n/en.yaml", yaml.Unmarshal)
	localizer.AddLocale("ru", "internal/i18n/ru.yaml", yaml.Unmarshal)

	// Run
	webappbot.Run(cfg, localizer)
}
