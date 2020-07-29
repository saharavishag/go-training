package main

// all keys and values should ve of the same type

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["black"] = "#000000"
	fmt.Println(colors)

	delete(colors, "black")
	fmt.Println(colors)

	printMap(colors)

	var colorsInit map[string]string
	fmt.Println(colorsInit)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("the color", color, "is represented by", hex)
	}
}
