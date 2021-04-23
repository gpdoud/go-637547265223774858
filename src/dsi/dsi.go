package dsi

import (
	"io/ioutil"
	"net/http"
)

// About ...
// Returns: nothing
func About() {
	println("About dsi")
}

// HTTPGet ..
func HTTPGet(url string) (s string, b string) {
	resp, _ := http.Get(url)
	s = resp.Status
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	b = string(body)
	return s, b
}
