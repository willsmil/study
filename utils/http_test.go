package utils

import (
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	url := "http://127.0.0.1:18008/internal/v1/gray-rule/hit"
	body := `{"device_id":"d8f36c8af3d9515f3064cff3ebc87381","ip":"","gray_rules":[{"rule_id":"1","business_id":"test"}]}`
	resp, err := PostForm(url, nil, body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", resp)
}
