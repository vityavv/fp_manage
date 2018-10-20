package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Games []struct{
	Name string
	Path string
}

func GetGames() (Games, error) {
	file, err := ioutil.ReadFile(os.Getenv("HOME") + "/.fpm_games")
	if err != nil {
		return Games{}, err
	}
	var games Games
	err = json.Unmarshal(file, &games)
	if err != nil {
		return Games{}, err
	}
	return games, nil
}