package main

import (
	"log"
	"os"

	"github.com/b4b4r07/go-crowi"
)

func main() {
	// dummy token
	client, err := crowi.NewClient("http://localhost:3000", "ywAVcbOsAKKwL7y8AkwXdxkLDO1YsqXwHl4oYYwYHMw=")
	if err != nil {
		panic(err)
	}

	path := "/user/john/memo"
	if len(os.Args[1:]) > 0 {
		path = os.Args[1]
	}

	resp, err := client.PagesCreate(path, "# this is a sample")
	if err != nil {
		panic(err)
	}

	if !resp.OK {
		log.Printf("%s\n", resp.Error)
	}

	resp, err = client.PagesUpdate(resp.Page.ID, "# this is a sample!!")
	if err != nil {
		panic(err)
	}

	if !resp.OK {
		log.Printf("%s\n", resp.Error)
	}

	resp2, err := client.AttachmentsAdd(resp.Page.ID, "./gopher.png")
	if err != nil {
		panic(err)
	}

	if !resp2.OK {
		log.Printf("%s\n", resp2.Error)
	}
}
