package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/integrii/flaggy"
)

var version = "unknown"
var singleFile string

func init() {
	flaggy.SetName("Test Program")
	flaggy.SetDescription("A little example program")
	flaggy.String(&singleFile, "s", "single", "A variable just for testing things!")
	flaggy.Parse()
}

func singleFileRename(str []string) {
	if len(str) < 2 {
		fmt.Printf("File rename error: Not enough arguments\r\n")
	}
	err := os.Rename(str[0], str[1])
	if err != nil {
		log.Fatalf("File rename error: %v\r\n", err)
	} else {
		fmt.Println(str[0], "=>", str[1], "Success!")
	}
}

func main() {
	// Parsing command-line arguments
	// Single file full path rename (short style)
	if len(singleFile) > 0 {
		singleFileRename(strings.Split(singleFile, "??"))
	}

}
