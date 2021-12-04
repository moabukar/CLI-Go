package main

import (
	"flag"
	"fmt"
	"os"
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

	if len(os.Args) < 2 {
		fmt.Println("exoected 'get' or 'add' subcommands")
		os.Exit()
	}

	// switch statement
	switch os.Args[1] {
	case "get":

	case "add":

	default:
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {

}
