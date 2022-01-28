package logservice

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLogHttpRequestOk(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	LogHttpRequest(200, "GET", "/")
	log.SetOutput(os.Stdout)
	if !strings.Contains(buf.String(), "[INFO] \033[33mOK\033[0m GET /") {
		t.Errorf("Incorrect log messsage received")
	}
}

func TestLogError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	LogError(400, "POST", "/", sql.ErrNoRows)
	log.SetOutput(os.Stdout)
	if !strings.Contains(buf.String(), fmt.Sprintf("[ERROR] \033[31mBad Request\033[0m POST / %s", sql.ErrNoRows)) {
		t.Errorf("Incorrect log messsage received")
	}
}

func TestLogInfo(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	LogInfo("Hello, World!")
	log.SetOutput(os.Stdout)
	if !strings.Contains(buf.String(), fmt.Sprintf("[INFO] Hello, World!")) {
		t.Errorf("Incorrect log messsage received")
	}
}
