package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	delete(colors, "red")

	fmt.Println("after delete red, ", colors)

	//this works as well: var colors2 map[string]string
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color of ", color, "has hex code ", hex)
	}
}
