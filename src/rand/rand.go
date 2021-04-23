package main

import (
	"dsi"
)

func main() {
	dsi.About()
	var status, body = dsi.HTTPGet("http://localhost:5000/api/users")
	println(status, body)
}
