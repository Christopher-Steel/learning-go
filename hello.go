package main

import "fmt"

func Yo(name string) string {
	return "Yo, yoyoyo, " + name + "!"
}

func main() {
	fmt.Println(Yo("banana"))
}
