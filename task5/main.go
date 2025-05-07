/*
Copyright © 2025 den-vasyliev <den.vasyliev@gmail.com>
*/
package main

import (
	"os"

	"github.com/den-vasyliev/kbot/cmd"
)

var (
	TeleToken string = os.Getenv("TELE_TOKEN")
)

func main() {
	cmd.Execute()
}
