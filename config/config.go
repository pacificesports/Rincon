package config

import (
	"os"
	"rincon/model"
)

var Service = model.Service{}

var Version = "2.1.1"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")
var JaegerPort = os.Getenv("JAEGER_PORT")

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
var StatusChannel = os.Getenv("STATUS_CHANNEL")

var StatusEmail = os.Getenv("STATUS_EMAIL")

var RegistryUpdateDelay = os.Getenv("REGISTRY_UPDATE_DELAY")
