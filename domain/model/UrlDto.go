package model

import (
	"cf-proposal/common/errormessages"
	"fmt"
)

type UrlDto struct {
	UrlID        int32
	LongUrl      string
	ShortUrl     string
	ExpirationDt string
}

func (dto UrlDto) ValidateCreate() error {

	if dto.LongUrl == "" {
		message := fmt.Sprintf(errormessages.INVALID_NULL_VALUE, "Long Url")
		return fmt.Errorf("%s", message)
	}
	return nil
}
