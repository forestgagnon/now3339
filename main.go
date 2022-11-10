package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const syntaxMsg = "error: syntax is now3339 [plus,minus] [duration] (Golang duration syntax)]"

func main() {
	if len(os.Args) == 1 {
		fmt.Println(time.Now().Format(time.RFC3339))
		os.Exit(0)
	}

	if len(os.Args) != 3 {
		log.Fatal(syntaxMsg)
	}

	var t time.Time
	dur := duration(os.Args[2])
	if os.Args[1] == "plus" {
		t = time.Now().Add(dur)
	} else if os.Args[1] == "minus" {
		t = time.Now().Add(-dur)
	} else {
		log.Fatal(syntaxMsg)
	}

	fmt.Println(t.Format(time.RFC3339))
}

func duration(input string) time.Duration {
	duration, err := time.ParseDuration(input)
	if err != nil {
		log.Fatalf("invalid Go duration syntax: %s", err.Error())
	}

	return duration
}
