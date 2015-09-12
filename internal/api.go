// {{{ Copyright (c) Paul R. Tagliamonte <paultag@dc.cant.vote>, 2015
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

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
	fmt.Printf("%s\n", resource)
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

// vim: foldmethod=marker
