package main

import (
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 {

		switch args[0] {
		case "add":
			Add(args)
		case "remove":
		case "update":
		case "list":
		case "mark":
		}

	}
}

func OpenFile() *os.File {
	file, err := os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Add(args []string) {

	file := OpenFile()

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

}

func Remove(args []string) {

}

func Update(args []string) {

}

func List(args []string) {

}

func Mark(args []string) {

}
