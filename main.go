package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Karitham/corde"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using env variable")
	}

	var command = corde.NewSlashCommand("random-song", "send random song")
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN not set")
	}
	appID := corde.SnowflakeFromString(os.Getenv("DISCORD_APP_ID"))
	if appID == 0 {
		log.Fatalln("DISCORD_APP_ID not set")
	}
	pk := os.Getenv("DISCORD_PUBLIC_KEY")
	if pk == "" {
		log.Fatalln("DISCORD_PUBLIC_KEY not set")
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("PORT environment variable not set, falling back to default 8080")
		port = 8080
	}

	m := corde.NewMux(pk, appID, token)
	m.SlashCommand("random-song", randomSongHandler)

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	log.Printf("serving on :%d\n", port)
	if err := m.ListenAndServe(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln(err)
	}

	songId, err := getRandomSongId()
	if err != nil {
		fmt.Println("Error in getting random song")
	}
	fmt.Println(getKosamegaPost(songId))
}

func randomSongHandler(ctx context.Context, w corde.ResponseWriter, _ *corde.Interaction[corde.SlashCommandInteractionData]) {
	songId, err := getRandomSongId()
	if err != nil {
		fmt.Println("Error in getting random song")
		w.Respond(corde.NewResp().Content("error getting random song").Ephemeral())
	}
	post := getKosamegaPost(songId)
	w.Respond(corde.NewResp().Content(post))
}
