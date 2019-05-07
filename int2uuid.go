package main

import (
	"fmt"
	"os"

	"github.com/pylebecq/int2uuid/uuid"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	uuid, err := uuid.StringToUUID(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(uuid.String())
}

func usage() {
	fmt.Printf("Usage: %s <int>\n", os.Args[0])
	fmt.Println("Arguments:")
	fmt.Println("  int: The integer to convert to a UUID")
}
