package helper

import (
	"testing"
	"time"
)

func TestUniqueShortUrl(t *testing.T) {

	longUrlOne := "https://www.youtube.com/watch?v=YS4e4q9oBaU"
	longUrlTwo := "https://www.youtube.com/watch?v=uBPXNREhZZw"

	shortUrlOne := getShortUrl(longUrlOne)
	shortUrlTwo := getShortUrl(longUrlTwo)

	if shortUrlOne == shortUrlTwo {
		t.Errorf("Hashes should not match")
	}
}

func TestNonDuplicatedShortUrl(t *testing.T) {
	longUrl := "https://www.youtube.com/watch?v=YS4e4q9oBaU"
	shortUrlOne := getShortUrl(longUrl)

	// Delay in requests
	time.Sleep(3 * time.Second)

	shortUrlTwo := getShortUrl(longUrl)

	if shortUrlOne != shortUrlTwo {
		t.Errorf("Hashes for the same url should match")
	}
}
