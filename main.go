package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	type GithubFile struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		DownloadUrl string `json:"download_url"`
	}
	fmt.Println("WELCOME TO MY FIRST GOLANG APP!!!")
	fmt.Println("Music Library :)")
	resp, err := http.Get("https://api.github.com/repos/Wyalii/music-files/contents/")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		var files []GithubFile

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &files)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(files)
		fmt.Println("Download URL specifically:", files[0].DownloadUrl)
	}

}
