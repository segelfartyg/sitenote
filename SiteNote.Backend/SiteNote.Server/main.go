package main

import (
	"notelad.com/server/db"
	"notelad.com/server/routing"
)

// SETTING UP WEB API.
func main() {
	db.Setup()
	routing.Setup()
}
