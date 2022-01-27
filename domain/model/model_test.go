package model

import (
	"cf-proposal/common/errormessages"
	"fmt"
	"testing"
)

func TestValidateCreateUrlFail(t *testing.T) {
	dto := UrlDto{
		LongUrl: "",
	}
	err := dto.ValidateCreate()

	if err.Error() != fmt.Sprintf(errormessages.INVALID_NULL_VALUE, "Long Url") {
		t.Errorf("Error message not matching, expected: %s, received: %s", fmt.Sprintf(errormessages.INVALID_NULL_VALUE, "Long Url"), err.Error())
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
