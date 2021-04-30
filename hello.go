package main

import "fmt"

func Yo(name string) string {
	const yoPrefix = "Yo, yoyoyo, "
	const yoSuffix = "!"
	if name == "" {
		name = "man"
	}
	return yoPrefix + name + yoSuffix
}

func main() {
	fmt.Println(Yo("banana"))
}
