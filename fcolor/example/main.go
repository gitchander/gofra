package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gitchander/gofra/fcolor"
)

func main() {

	type NameColor struct {
		Name  string
		Color fcolor.RGB
	}

	colors := []NameColor{
		NameColor{Name: "Black", Color: fcolor.RGB{R: 0, G: 0, B: 0}},
		NameColor{Name: "White", Color: fcolor.RGB{R: 1, G: 1, B: 1}},
		NameColor{Name: "Gray", Color: fcolor.RGB{R: 0.5, G: 0.5, B: 0.5}},
		NameColor{Name: "Red", Color: fcolor.RGB{R: 1, G: 0, B: 0}},
		NameColor{Name: "Green", Color: fcolor.RGB{R: 0, G: 0.5, B: 0}},
		NameColor{Name: "Blue", Color: fcolor.RGB{R: 0, G: 0, B: 1}},
		NameColor{Name: "Yellow", Color: fcolor.RGB{R: 1, G: 1, B: 0}},
		NameColor{Name: "Aqua", Color: fcolor.RGB{R: 0, G: 1, B: 1}},
		NameColor{Name: "Magenta", Color: fcolor.RGB{R: 1, G: 0, B: 1}},
	}

	for _, c := range colors {
		data, err := json.Marshal(&c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
