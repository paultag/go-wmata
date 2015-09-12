package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var apiKey string
var apiHost = "https://api.wmata.com"

func SetAPIKey(newApiKey string) {
	apiKey = newApiKey
}

//
func Get(resource string, params map[string]string, target interface{}) error {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	values.Set("api_key", apiKey)
	resource = fmt.Sprintf("%s/%s?%s", apiHost, resource, values.Encode())
	// fmt.Printf("%s\n", resource)
	resp, err := http.Get(resource)
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: Non-200 status code: %d", resp.StatusCode)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
