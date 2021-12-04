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
	case "get": // if get command is used
		HandleGet(getCmd, getAll, getID)
	case "add": // if add command is used
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default: // if input is not clear
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *id = "" {
		fmt.Println("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		video := getVideos()

		fmt.Printf()("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageURl, video.Description)
		}
	return
	
	}
	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n",video.Id, video.Title, video.Url, video.Imageurl,video.Description)

			}
		}
} 

}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	ValidateVideo(addCmd, id,title,url, imageUrl, description)

	video := video{
		Id: *id,
		Title: *title,
		Description: *description,
		Imageurl: *imageUrl, 
		Url: *url, 
	}

	videos := getVideos()
	videos = append(videos,video)

	saveVideos(videos)
}

func ValidateVideo(addCmd *flag.FlagSet,id *string, title *string, url *string, imageUrl *string, description *string ){
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}