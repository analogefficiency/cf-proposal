package helper

import (
	"testing"
	"time"
)

func TestErrorOk(t *testing.T) {
	m := "This is a test error"
	e := &CustomError{Message: m}
	if e.Error() != m {
		t.Errorf("Error message not matching")
	}
}

func TestNonDuplicatedShortUrl(t *testing.T) {
	longUrl := "https://www.youtube.com/watch?v=YS4e4q9oBaU"
	shortUrlOne := GetShortUrl(longUrl)

	// Delay in requests
	time.Sleep(3 * time.Second)

	shortUrlTwo := GetShortUrl(longUrl)

	if shortUrlOne != shortUrlTwo {
		t.Errorf("Hashes for the same url should match")
	}
}

func TestUniqueShortUrl(t *testing.T) {

	longUrlOne := "https://www.youtube.com/watch?v=YS4e4q9oBaU"
	longUrlTwo := "https://www.youtube.com/watch?v=uBPXNREhZZw"

	shortUrlOne := GetShortUrl(longUrlOne)
	shortUrlTwo := GetShortUrl(longUrlTwo)

	if shortUrlOne == shortUrlTwo {
		t.Errorf("Hashes should not match")
	}
}
