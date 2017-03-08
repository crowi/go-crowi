package main

import (
	"log"

	"github.com/b4b4r07/go-crowi"
)

func main() {
	// dummy token
	client, err := crowi.NewClient("ywAVcbOsAKKwL7y8AkwXdxkLDO1YsqXwHl4oYYwYHMw=", "http://localhost:3000")
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
