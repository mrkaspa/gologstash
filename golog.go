package golog

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Client that'll send the posts
type Client struct {
	URL      string
	User     string
	Password string
	Host     string
}

//PostLog a string
func (c Client) PostLog(log string) (int, error) {
	return c.PostMap(map[string]string{"message": log})
}

//PostMap a map as message json
func (c Client) PostMap(data map[string]string) (int, error) {
	postData := make(map[string]interface{})
	if message, ok := data["message"]; ok {
		postData["message"] = message
		delete(data, "message")
	}
	postData["add_field"] = data
	postData["host"] = c.Host
	jsonString, err := json.Marshal(postData)
	if err != nil {
		return 0, err
	}
	resp, err := c.postJSON(string(jsonString))
	return resp.StatusCode, err
}

func (c Client) postJSON(json string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", c.URL, strings.NewReader(json))
	req.Header.Add("Content-Type", "application/json")
	if c.User != "" {
		req.SetBasicAuth(c.User, c.Password)
	}
	return client.Do(req)
}
