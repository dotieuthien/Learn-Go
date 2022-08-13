package main

import "fmt"

func main() {
	// key type string and value type string
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"

	delete_color(colors, "red")
	iter_over_map(colors)
}

func iter_over_map(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, hex)
	}
}

func delete_color(m map[string]string, k string) map[string]string {
	delete(m, k)
	return m
}