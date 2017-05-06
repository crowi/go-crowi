package main

import (
	"context"
	"fmt"
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

	var (
		path = fmt.Sprintf("/user/%s/go-crowi-test-%d", os.Getenv("USER"), time.Now().UnixNano())
		body = "# this is a sample\n\ntest"
	)

	res1, err := client.Pages.Create(ctx, path, body)
	if err != nil {
		panic(err)
	}

	res2, err := client.Attachments.Add(ctx, res1.Page.ID, "_example/attachments/sample.png")
	if err != nil {
		panic(err)
	}
	pp.Println(res2)

	res3, err := client.Attachments.List(ctx, res1.Page.ID)
	if err != nil {
		panic(err)
	}

	// body = fmt.Sprintf("![](%s)", res3.URL)
	_, err = client.Pages.Update(ctx, res1.Page.ID, body)
	if err != nil {
		panic(err)
	}

	pp.Println(res3)
}
