package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Note struct {
	Title       string   `json:"Title"`
	Date        string   `json:"Date"`
	Description string   `json:"Description"`
	Complete    bool     `json:"Complete"`
	Tags        []string `json:"Tags"`
}

func (n Note) String() string {
	format := "\nTitle: %v \nDate: %v\nDescription: %v\nComplete: %v\nTags:%v\n"
	return fmt.Sprintf(format,
		n.Title, n.Date, n.Description, n.Complete, n.Tags)
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

func Add(path string, notes *map[int]Note, newnote Note) {
	l := len(*notes)
	(*notes)[l+1] = newnote
}

func Write(notes *map[int]Note) {
	wr, err := json.Marshal(notes)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("output.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	_, err = f.Write(wr)
	if err != nil {
		log.Fatalln(err)
	}
}

func Delete(notes *map[int]Note, n int) {
}
