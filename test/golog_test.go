package test

import (
	"fmt"
	"testing"

	"github.com/mrkaspa/gologstash"
)

var client = golog.Client{URL: "http://localhost:9090", Host: "localhost"}

func TestPostLog(t *testing.T) {
	if status, err := client.PostLog("demo"); err != nil {
		t.Fatal("Error posting the message", err)
	} else {
		fmt.Printf("status >> %d", status)
		if status != 200 {
			t.Errorf("the status is not ok >> %d", status)
		}
	}
}

func TestPostMap(t *testing.T) {
	data := map[string]string{"message": "demo message", "add": "I'm additional info"}
	if status, err := client.PostMap(data); err != nil {
		t.Fatal("Error posting the message", err)
	} else {
		fmt.Printf("status >> %d", status)
		if status != 200 {
			t.Errorf("the status is not ok >> %d", status)
		}
	}
}
