go-crowi
========

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godoc]

[license]: https://github.com/b4b4r07/go-crowi/blob/master/LICENSE
[godoc]: http://godoc.org/github.com/b4b4r07/go-crowi

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
