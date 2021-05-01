package main

import "fmt"

type Language struct {
	prefix      string
	suffix      string
	defaultName string
}

var langs = map[string]Language{
	"english": Language{
		"Yo, yoyoyo, ",
		"!",
		"man",
	},
	"spanish": Language{
		"Jo, jojojo, ",
		"!",
		"hombre",
	},
}

func Yo(name, language string) string {
	if language == "" {
		language = "english"
	}
	if name == "" {
		name = langs[language].defaultName
	}
	return langs[language].prefix + name + langs[language].suffix
}

func main() {
	fmt.Println(Yo("banana", ""))
}
