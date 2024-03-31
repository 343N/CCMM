package main

import (
	"fmt"
	"net/http"
)

var CORTEX_COMMAND_ID = "cccp"

func GetModsFromRemote() ([]ModData, error) {
	url := fmt.Sprintf("https://mod.io/games/cccp/mods?api_key=%s", GetMMData().ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		logger.Println("There was an issue getting mod info from mod.io!")
		return nil, err
	}

	var body []byte
	_, err = resp.Body.Read(body)
	resp

	if err != nil {
		logger.Println("There was an issue reading request data from mod.io!")
		return nil, err
	}

}

func DownloadMod(d ModData) {

}
