package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Note struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Date        string `json:"Date"`
	Description string `json:"Description"`
	Complete    bool   `json:"Complete"`
}

func Read(path string) map[int]Note {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	var notes map[int]Note
	err = json.Unmarshal(file, &notes)
	if err != nil {
		log.Fatalln(err)
	}

	return notes
}

func main() {
	/*
		n := map[int]note{
			1: {
				Id:          "1",
				Title:       "Title",
				Date:        "Today",
				Description: "Test,
				Complete:    false,
			},
			2: {
				Id:          "2",
				Title:       "Title",
				Date:        "Today",
				Description: "Test, test",
				Complete:    false,
			},
		}
	*/
	path := "tasks.json"
	notes := Read(path)

	for i, v := range notes {
		fmt.Printf("i: %d -> %v\n", i, v)
	}
}
