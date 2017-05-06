package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/crowi/go-crowi"
	"github.com/k0kubun/pp"
)

func main() {
	client, err := crowi.NewClient("http://localhost:3000", os.Getenv("CROWI_ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var (
		path = fmt.Sprintf("/user/%s/go-crowi-test-%d", os.Getenv("USER"), time.Now().UnixNano())
		body = "# this is a sample\n\ntest"
	)

	res, err := client.Pages.Create(ctx, path, body)
	if err != nil {
		panic(err)
	}

	if !res.OK {
		log.Printf("[ERROR] %s", res.Error)
		os.Exit(1)
	}

	pp.Println(res)
}
