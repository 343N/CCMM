package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ModioModRequest struct {
	Mods []ModioModData `json:"data"`
}

type RequestResultInfo struct {
	Count  int `json:"result_count"`
	Offset int `json:"result_offset"`
	Limit  int `json:"result_limit"`
	Total  int `json:"result_total"`
}

type ModioModData struct {
	Id      int64
	Visible int
	Status  int
	Name    string
	Name_id string
	Summary string
	Thumb   string
	Logo    struct {
		Original string
	}
	Modfile struct {
		Id       int64
		Filesize int64
		Filename string
		Version  string
		Download struct {
			Binary_url   string
			Date_expires int64
		}
	}
	Submitted_by struct {
		Id          int64
		name_id     string
		Username    string
		Profile_url string
	}
}

var CORTEX_COMMAND_MOD_ID = "508"
var CORTEX_COMMAND_NAME_ID = "cccp"

func GetModsFromRemote(offset int) ([]MMModData, *RequestResultInfo, error) {
	url := fmt.Sprintf("https://u-%s.modapi.io/v1/games/%s/mods?api_key=%s",
		GetMMData().Api.UserId,
		CORTEX_COMMAND_MOD_ID,
		GetMMData().Api.Key,
	)

	logger.Printf("Grabbing mod data from %s...", url)

	resp, err := http.Get(url)
	if err != nil {
		logger.Println("There was an issue getting mod info from mod.io!")
		return nil, nil, err
	}

	// 1MB limit
	body, err := io.ReadAll(resp.Body)
	// remove null chars lel
	body = bytes.Replace(body, []byte("\x00"), []byte(""), -1)

	if err != nil {
		logger.Println("There was an issue reading request data from mod.io!")
		return nil, nil, err
	}

	var reqData = &ModioModRequest{}
	var reqResultInfo = &RequestResultInfo{}
	err = json.Unmarshal(body, reqData)
	err = json.Unmarshal(body, reqResultInfo)
	if err != nil {
		logger.Println("There was an issue unmarshalling request data from mod.io!")
		logger.Printf("Body size: %d", len(string(body)))
		logger.Printf("Error: %v", err)
		return nil, nil, err
	}
	logger.Println("Retrieved mods from remote! Count: %s", len(reqData.Mods))

	mm_mods := []MMModData{}

	for _, val := range reqData.Mods {
		mm_mods = append(mm_mods, MMModData{
			Name:        val.Name,
			Id:          val.Id,
			Author:      val.Submitted_by.Username,
			DownloadUrl: val.Modfile.Download.Binary_url,
			FileName:    val.Modfile.Filename,
			FileSize:    int(val.Modfile.Filesize),
			ThumbUrl:    val.Logo.Original,
			ModVersion:  val.Modfile.Version,
		})
	}

	logger.Printf("Debug: %+v", reqData.Mods[0])
	logger.Printf("Debug: %+v", mm_mods[0])

	return mm_mods, reqResultInfo, nil

}

func DownloadMod(d *MMModData) {

}
