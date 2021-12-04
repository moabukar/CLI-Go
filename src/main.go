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
	addID := addCmd.String("id", "", "Video ID")
	addTitle := addCmd.String("title", "", "Video Title")
	addUrl := addCmd.String("Url", "", "Video URL")
	addImageUrl := addCmd.String("iamge URL", "", "video image URL")
	addDesc := addCmd.String("desc", "", "video description")
}
