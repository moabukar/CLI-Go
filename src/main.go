package main

import (
	"flag"
)

func main() {
	// subcommand = "videos get"
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// subcommand for video get all & ID
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	//
	add := addCmd.String()

}
