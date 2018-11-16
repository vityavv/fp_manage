package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Game struct {
	Name string
	Path string
}

type Games []Game

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

func AddGame(name, path string) error {
	currgames, err := GetGames()
	if err != nil {
		return err
	}
	currgames = append(currgames, Game{
		Name: name,
		Path: path,
	})
	return writeJSON(currgames)
}

func DeleteGame(id int) error {
	currgames, err := GetGames()
	if err != nil {
		return err
	}
	//trick to remove
	currgames = append(currgames[:id], currgames[id+1:]...)
	return writeJSON(currgames)
}
	
func writeJSON(data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(os.Getenv("HOME") + "/.fpm_games", json, 0600)
	if err != nil {
		return err
	}
	return nil
}
