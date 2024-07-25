package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/oauth2/clientcredentials"
)

type ApiResponse struct {
	Limit  int    `json:"limit"`
	Items  []Item `json:"items"`
	Offset int    `json:"offset"`
}

type Item struct {
	Track Track `json:"track"`
}

type Track struct {
	Album Album  `json:"album"`
	Href  string `json:"href"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}

type Album struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

func fetchAllData(baseURL *url.URL, client *http.Client) ([]Item, error) {
	var allData []Item
	offset := 0
	limit := 100
	fields := "offset,limit,items(track(id,name,href,album(name,href)))"

	for {
		params := url.Values{}
		params.Add("offset", strconv.Itoa(offset))
		params.Add("limit", strconv.Itoa(limit))
		params.Add("fields", fields)
		baseURL.RawQuery = params.Encode()

		resp, err := client.Get(baseURL.String())
		if err != nil {
			fmt.Println("Error making request:", err)
			return nil, err
		}
		//fmt.Println(resp)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}

		var apiResponse ApiResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			return nil, err
		}

		if len(apiResponse.Items) == 0 {
			break
		}

		allData = append(allData, apiResponse.Items...)
		offset += limit
	}

	return allData, nil
}

func getOpenSpotifyURL(id string) string {
	return fmt.Sprintf("https://open.spotify.com/track/%s", id)
}

func getRandomSongId() (string, error) {
	config := clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     "https://accounts.spotify.com/api/token",
	}

	_, err := config.Token(context.Background())
	if err != nil {
		fmt.Printf("Error retrieving token: %v\n", err)
		return "", err
	}
	client := config.Client(context.Background())

	playlistId := os.Getenv("SPOTIFY_PLAYLIST_ID")
	baseURL, err := url.Parse(fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistId))
	if err != nil {
		fmt.Println("URL parse error:", err)
		return "", err
	}

	allData, err := fetchAllData(baseURL, client)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return "", err
	}

	fmt.Printf("Total data count: %d\n", len(allData))
	return randomChoice(allData).Track.Id, nil
}
