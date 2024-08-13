package main

import (
	"math/rand"
	"strings"
	"time"
)

type RandSource interface {
	Intn(n int) int
}

type DefaultRandSource struct{}

func (d DefaultRandSource) Intn(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

func randomChoice[T any](items []T, r RandSource) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[r.Intn(len(items))]
}

func randomInt(n int, r RandSource) int {
	return r.Intn(n)
}

func getKosamegaPost(songId string, r RandSource) string {
	var partial []string
	start := []string{"冷静に聴き直し、", "久しぶりに聴いてみたら", "思い出して聴いてみたら"}
	partial = append(partial, randomChoice(start, r))
	majide := []string{"マジで", "本当に", "かなり", "めっちゃ", "世界一"}
	partial = append(partial, randomChoice(majide, r))
	sugo := []string{"凄", "ヤバ", "かっこよ", "良"}
	partial = append(partial, randomChoice(sugo, r))

	if randomInt(2, r)%2 == 0 {
		kute := []string{"くね？？？？？", "すぎ？？？"}
		partial = append(partial, randomChoice(kute, r))
	} else {
		kute := []string{"すぎて", "くて"}
		partial = append(partial, randomChoice(kute, r))
		kandou := []string{"感動してる", "ぶちあがってる", "くらってる", "頭抱えてる", "エグい", "狂ってきた", "大好き", "最高"}
		partial = append(partial, randomChoice(kandou, r))
	}
	partial = append(partial, "\n"+getOpenSpotifyURL(songId))

	result := strings.Join(partial, "")
	return result
}
