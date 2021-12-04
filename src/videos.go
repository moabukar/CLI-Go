package main

import (
	"encoding/json"
	"io/ioutil"
)
type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

func getVideos()(video []video){

	fileBytes, err !: ioutil.ReadFile("./videos.json")

	if err != nil {    //error handling here
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err!= nil {
		panic(err)
	}

	return videos
}

func saveVideos()(video []video){

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}
	
	err = ioutil.WriteFile("./videos.json",videoBytes,0644)
	if err != nil {
		panic(err)
	}
}