package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	fmt.Println(colors)

	printMap(colors)

	colors["white"] = "#FFFFFF"
	fmt.Println(colors)

	delete(colors, "red")
	fmt.Println(colors)
}
