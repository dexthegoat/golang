package main

import "fmt"

// map: reference type
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#451225f",
		"white": "#ffffff",
	}

	// var colors2 map[string]string
	colors2 := make(map[int]string)
	// struct不能这样直接增删属性
	colors2[10] = "#ffffff"
	delete(colors2, 10)

	fmt.Println(colors)
	fmt.Println(colors2)

	// 遍历map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + ": " + hex)
	}
}
