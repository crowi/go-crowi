package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/crowi/go-crowi"
	"github.com/k0kubun/pp"
)

func main() {
	config := crowi.Config{
		URL:   "http://localhost:3000",
		Token: os.Getenv("CROWI_ACCESS_TOKEN"),
	}
	client, err := crowi.NewClient(config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.Pages.List(ctx, "", os.Getenv("USER"), &crowi.PagesListOptions{
		ListOptions: crowi.ListOptions{Pagenation: true},
	})
	if err != nil {
		panic(err)
	}

	if !res.OK {
		log.Printf("[ERROR] %s", res.Error)
		os.Exit(1)
	}

	pp.Println(res)
}
