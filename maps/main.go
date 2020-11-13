// It's like Hash in Ruby, Object in JavaScript and Dictionary in Python
// In Go, it works like an array to assign values
package main

import (
	"fmt"
)

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"blue": "#0000ff",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// Append values
	colors["black"] = "#000000"

	// Delete values
	delete(colors, "green")

	printMap(colors)
}

func printMap(c map[string]string) {
	// Iterate values
	for color, value := range c {
		fmt.Println("Color ", color, " in hex = ", value)
	}
}