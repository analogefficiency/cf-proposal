package helper

import (
	"encoding/hex"
	"hash/fnv"
)

func getShortUrl(longUrl string) string {

	hash := fnv.New32a()
	hash.Write([]byte(longUrl))
	return hex.EncodeToString(hash.Sum(nil))
}
