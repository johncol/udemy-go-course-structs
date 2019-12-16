package main

import "fmt"

func mapsSample() {
	colors := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	colors["red"] = "#ff0000"
	fmt.Println(colors)

	delete(colors, "white")
	fmt.Println(colors)

	for color, hex := range colors {
		fmt.Println("Hex value for color", color, "is", hex)
	}
}
