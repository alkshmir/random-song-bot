package main

import (
	"fmt"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/amatsagu/tempest"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Error loading .env file, using env variable")
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		slog.Info("PORT environment variable not set, falling back to default 8080")
		port = 8080
	}

	slog.Info("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: os.Getenv("DISCORD_PUBLIC_KEY"),
		Rest:      tempest.NewRestClient(os.Getenv("DISCORD_BOT_TOKEN")),
	})

	client.RegisterCommand(GetRandomSong)
	err = client.SyncCommands([]tempest.Snowflake{}, nil, false)
	if err != nil {
		slog.Error("failed to sync local commands storage with Discord API", err)
	}
	http.HandleFunc("POST /discord/callback", client.HandleDiscordRequest)

	slog.Info(fmt.Sprintf("Serving application at: :%d/discord/callback", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		slog.Error("Fatal: ", err)
	}

}

var GetRandomSong tempest.Command = tempest.Command{
	Name:          "random-song",
	Description:   "send random song",
	AvailableInDM: true,
	SlashCommandHandler: func(itx *tempest.CommandInteraction) {
		songId, err := getRandomSongId()
		if err != nil {
			slog.Error("Error in getting random song")
			itx.SendLinearReply("error getting random song", true)
			return
		}
		post := getKosamegaPost(songId)
		itx.SendLinearReply(post, false)
		//itx.SendLinearReply(fmt.Sprintf("Result: %d", af+bf), false)
	},
}
