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
	"strings"
	"time"

	"math/rand"

	"github.com/joho/godotenv"
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

func randomChoice[T any](items []T) T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if len(items) == 0 {
		// スライスが空の場合はゼロ値を返す
		var zero T
		return zero
	}

	randomIndex := r.Intn(len(items))

	return items[randomIndex]
}

func randomInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func getOpenSpotifyURL(id string) string {
	return fmt.Sprintf("https://open.spotify.com/track/%s", id)
}

func getKosamegaPost(songId string) string {
	var partial []string
	start := []string{"冷静に聞き直すと", "久しぶりに", "思い出して"}
	partial = append(partial, randomChoice(start))
	partial = append(partial, "聞いてみたら")
	majide := []string{"マジで", "本当に", "かなり", "めっちゃ", "世界一"}
	partial = append(partial, randomChoice(majide))
	sugo := []string{"凄", "ヤバ", "かっこよ", "良"}
	partial = append(partial, randomChoice(sugo))

	if randomInt(2)%2 == 0 {
		kute := []string{"くね？？？？？", "すぎ？？？"}
		partial = append(partial, randomChoice(kute))
	} else {
		kute := []string{"すぎて", "くて"}
		partial = append(partial, randomChoice(kute))
		kandou := []string{"感動してる", "ぶちあがってる", "くらってる", "頭抱えてる", "エグい", "狂ってきた", "大好き", "最高"}
		partial = append(partial, randomChoice(kandou))
	}
	partial = append(partial, "\n"+getOpenSpotifyURL(songId))

	result := strings.Join(partial, "")
	return result
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using env variable")
	}

	config := clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     "https://accounts.spotify.com/api/token",
	}

	config.Token(context.Background())
	if err != nil {
		fmt.Printf("Error retrieving token: %v\n", err)
		return
	}
	client := config.Client(context.Background())

	playlistId := os.Getenv("SPOTIFY_PLAYLIST_ID")
	baseURL, err := url.Parse(fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistId))
	if err != nil {
		fmt.Println("URL parse error:", err)
		return
	}

	allData, err := fetchAllData(baseURL, client)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	fmt.Printf("Total data count: %d\n", len(allData))
	songId := randomChoice(allData).Track.Id
	//songUrl := getOpenSpotifyURL(songId)
	fmt.Println(getKosamegaPost(songId))
}
