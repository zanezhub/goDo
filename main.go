package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	path := "tasks.json"
	notes := Read(path)

	/*
			for i, v := range notes {
				fmt.Print(i)
				fmt.Println(v)
			}

		ADD()
		newnote := Note{
			Title:       "Test title",
			Date:        "12 NOV 2023",
			Description: "TTestingTestingTestingTestingesting",
			Complete:    false,
			Tags:        []string{"test", "tarea"},
		}
	*/
	p := tea.NewProgram(initialModel(&notes))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
