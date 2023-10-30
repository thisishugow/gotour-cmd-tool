package main

import (
	"log"

	"yuwang.idv/go-tour/cmd-tool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
