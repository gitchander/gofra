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
		NameColor{Name: "Black", Color: fcolor.RGB{0, 0, 0}},
		NameColor{Name: "White", Color: fcolor.RGB{1, 1, 1}},
		NameColor{Name: "Gray", Color: fcolor.RGB{0.5, 0.5, 0.5}},
		NameColor{Name: "Red", Color: fcolor.RGB{1, 0, 0}},
		NameColor{Name: "Green", Color: fcolor.RGB{0, 0.5, 0}},
		NameColor{Name: "Blue", Color: fcolor.RGB{0, 0, 1}},
		NameColor{Name: "Yellow", Color: fcolor.RGB{1, 1, 0}},
		NameColor{Name: "Aqua", Color: fcolor.RGB{0, 1, 1}},
		NameColor{Name: "Magenta", Color: fcolor.RGB{1, 0, 1}},
	}

	for _, c := range colors {
		data, err := json.Marshal(&c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
