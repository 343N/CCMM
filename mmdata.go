package main

import (
	"bytes"
	"encoding/gob"
	"os"
)

type MMData struct {
	GamePath   string
	ModDir     string
	ApiKey     string
	LoadedMods []ModData
}

type ModData struct {
	Name       string
	Version    string
	Id         string
	FolderName string
	Thumbnail  []byte
	Downloaded bool
}

var Data *MMData = &MMData{}

func LoadMMData(filepath string) {
	if !fileExists(filepath) {
		logger.Printf("Data file at %s does not exist. Creating new data file!\n", filepath)
		SaveMMData(filepath)
		return
	}
	b, err := os.ReadFile(filepath)
	if err != nil {
		logger.Fatalf("Failed to read data from file at %s! Exiting!\n", filepath)
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	err = dec.Decode(Data)
	if err != nil {
		logger.Fatalf("Failed to decode data from file at %s! Exiting!\n", filepath)
	}
}

func GetMMData() *MMData {
	return Data
}

func SaveMMData(filepath string) error {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(Data)

	if err != nil {
		logger.Printf("Failed to save data!")
		return err
	}

	os.WriteFile(filepath, buffer.Bytes(), 0666)
	logger.Printf("Saved MM data to %s!\n", filepath)

	return nil
}
