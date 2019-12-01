package main

import (
	"os"
	"time"
)

var discordClient = discordConfig{
	Boss:      os.Getenv("DISCORD_BOSS"),
	Token:     os.Getenv("DISCORD_TOKEN"),
	AvatarURL: "https://img1.ak.crunchyroll.com/i/spire4/e3b3bae63ad7e179ef6106689cd3b1901553253026_large.png",
	Debug:     false,
}

func init() {
	requireEnvVars("DATABASE_HOST", "DATABASE_PORT", "DATABASE_DB", "DISCORD_BOSS", "DISCORD_TOKEN")
}

func main() {
	defer webServer()
	dbConn()

	go discordStart(&discordClient)
	go tickerHelper(10*time.Minute, rssReader, false)
	go tickerHelper(10*time.Hour, maintainAnimeList, true)
}
