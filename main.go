package main

import (
	"fmt"
	"flag"
	"github.com/satori/go.uuid"
	"github.com/atotto/clipboard"
)

func main() {
	uuid := uuid.NewV4()
	short := flag.Bool("s", false, "Print only uuid. Perfect for `uuid -s | pbcopy`")
	copy := flag.Bool("c", false, "Copy uuid to clipboard")
	flag.Parse()
	if *copy {
		clipboard.WriteAll(uuid.String())
	} else if *short {
		fmt.Printf("%v", uuid.String())
	} else {
		fmt.Printf("Your uuid is: %v\n", uuid.String())
	}
}
