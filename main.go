package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using env variable")
	}

	songId, err := getRandomSongId()
	if err != nil {
		fmt.Println("Error in getting random song")
	}
	fmt.Println(getKosamegaPost(songId))
}
