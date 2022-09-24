package main

import (
	"sesi6-gin/httpserver"
)

func main() {
	app := httpserver.CreateRouter()
	app.Run(":4000")
}
