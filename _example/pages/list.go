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
	client, err := crowi.NewClient("http://localhost:3000", os.Getenv("CROWI_ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.PagesList(ctx, "/user/john", "")
	if err != nil {
		panic(err)
	}

	if !res.OK {
		log.Printf("[ERROR] %s", res.Error)
		os.Exit(1)
	}

	pp.Println(res)
}
