package main

import (
	"testing"
)

type MockRandSource struct{}

func (m MockRandSource) Intn(n int) int {
	return 0 // 常に最初の要素を返す
}

func Test_getKosamegaPost(t *testing.T) {
	mockRand := MockRandSource{}

	expected := "冷静に聴き直し、マジで凄くね？？？？？\n" + getOpenSpotifyURL("songID")
	result := getKosamegaPost("songID", mockRand)

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
