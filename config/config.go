package config

import (
	"os"
	"rincon/model"
	"strings"
)

var Service = model.Service{
	Name:        os.Getenv("SERVICE_NAME"),
	StatusEmail: os.Getenv("STATUS_EMAIL"),
	URL:         "http://" + strings.ToLower(os.Getenv("SERVICE_NAME")) + ":" + Port,
	Version:     Version,
}

var Version = "2.1.6"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var PostgresHost = os.Getenv("POSTGRES_HOST")
var PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
var PostgresUser = os.Getenv("POSTGRES_USER")
var PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
var PostgresPort = os.Getenv("POSTGRES_PORT")

var EmailAddress = os.Getenv("EMAIL_ADDRESS")
var EmailPassword = os.Getenv("EMAIL_PASSWORD")

var DiscordToken = os.Getenv("DISCORD_TOKEN")
var DiscordGuild = os.Getenv("DISCORD_GUILD")
var DiscordChannel = os.Getenv("DISCORD_CHANNEL")

var StatusEmail = os.Getenv("STATUS_EMAIL")

var RegistryUpdateDelay = os.Getenv("REGISTRY_UPDATE_DELAY")
