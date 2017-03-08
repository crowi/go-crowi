go-crowi
========

A Go client for Crowi APIs

## Installation

```console
$ go get github.com/b4b4r07/go-crowi
```

## Example

```go
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
```

## License

MIT

## Author

b4b4r07
