package main

import "fmt"

func Yo(name string) string {
	const yoPrefix = "Yo, yoyoyo, "
	const yoSuffix = "!"
	return yoPrefix + name + yoSuffix
}

func main() {
	fmt.Println(Yo("banana"))
}
