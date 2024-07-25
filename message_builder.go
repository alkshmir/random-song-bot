package main

import (
	"math/rand"
	"strings"
	"time"
)

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
