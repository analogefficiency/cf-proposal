package model

import (
	"cf-proposal/common/messages"
	"fmt"
	"net/url"
)

type UrlDto struct {
	UrlID        int32
	LongUrl      string
	ShortUrl     string
	ExpirationDt string
}

func (dto UrlDto) ValidateCreate() error {

	if dto.LongUrl == "" {
		message := fmt.Sprintf(messages.INVALID_NULL_VALUE, "Long Url")
		return fmt.Errorf("%s", message)
	}

	_, err := url.ParseRequestURI(dto.LongUrl)
	if err != nil {
		message := fmt.Sprintf(messages.INVALID_URL, dto.LongUrl)
		return fmt.Errorf("%s", message)
	}
	return nil
}
