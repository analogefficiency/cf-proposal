package model

import (
	"cf-proposal/common/messages"
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type UrlDto struct {
	UrlID        int32
	LongUrl      string
	ShortUrl     string
	ExpirationDt string
}

func (dto UrlDto) ValidateCreate() error {

	var findings error
	if dto.LongUrl == "" {
		message := fmt.Sprintf(messages.INVALID_NULL_VALUE, "Long Url")
		findings = fmt.Errorf("%s", message)
	} else {
		_, err := url.ParseRequestURI(dto.LongUrl)
		if err != nil {
			message := fmt.Sprintf(messages.INVALID_URL, dto.LongUrl)
			findings = fmt.Errorf("%s", message)
		}
	}

	if dto.ExpirationDt != "" {
		_, err := time.Parse("2006-01-02 15:04:05", dto.ExpirationDt)
		if err != nil {
			message := fmt.Sprintf(messages.INVALID_DATETIME, dto.ExpirationDt)
			if findings != nil {
				findings = errors.Wrap(findings, message)
			} else {
				findings = fmt.Errorf("%s", message)
			}
		}
	}

	return findings
}
