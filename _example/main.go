package main

import (
	"log"

	"github.com/b4b4r07/crowi-go"
)

func main() {
	client, err := crowi.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CreatePage("/user/john/memo", "# this is a sample")
	if err != nil {
		log.Fatal(err)
	}

	if !resp.OK {
		log.Printf("%s\n", resp.Error)
	}
}
