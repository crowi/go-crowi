go-crowi
========

A Go client for Crowi APIs

## Example

```go
client, err := crowi.NewClient("http://localhost:3000", "abcdefghijklmnopqrstuvwxyz0123456789=")
if err != nil {
	panic(err)
}

res, err := client.PagesCreate("/user/john/memo", "# this is a sample")
if err != nil {
	panic(err)
}

if !res.OK {
	log.Printf("[ERROR] %s", res.Error)
}
```

## License

MIT

## Author

b4b4r07
