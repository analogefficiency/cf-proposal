package model

import (
	"cf-proposal/common/messages"
	"fmt"
	"testing"
)

func TestValidateCreateUrlIsEmpty(t *testing.T) {
	dto := UrlDto{
		LongUrl: "",
	}
	err := dto.ValidateCreate()

	if err.Error() != fmt.Sprintf(messages.INVALID_NULL_VALUE, "Long Url") {
		t.Errorf("Error message not matching, expected: %s, received: %s", fmt.Sprintf(messages.INVALID_NULL_VALUE, "Long Url"), err.Error())
	}
}

func TestValidateCreateUrlIsInvalid(t *testing.T) {
	dto := UrlDto{
		LongUrl: "http///youtube.com",
	}
	err := dto.ValidateCreate()

	if err.Error() != fmt.Sprintf(messages.INVALID_URL, dto.LongUrl) {
		t.Errorf("Error message not matching, expected: %s, received: %s", fmt.Sprintf(messages.INVALID_NULL_VALUE, "Long Url"), err.Error())
	}
}

func TestValidateCreateUrlOk(t *testing.T) {
	dto := UrlDto{
		LongUrl: "https://hawaiinewnow.com",
	}
	err := dto.ValidateCreate()

	if err != nil {
		t.Error("Error not expected")
	}
}
