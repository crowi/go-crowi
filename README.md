go-crowi [![][travis-badge]][travis-link] [![][license-badge]][license-link] [![][godoc-badge]][godoc-link]
========

A Go client for Crowi APIs

## Getting

```console
$ go get github.com/b4b4r07/go-crowi
```

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

[travis-badge]: http://img.shields.io/travis/b4b4r07/go-crowi.svg?style=flat-square
[travis-link]: https://travis-ci.org/b4b4r07/go-crowi

[license-badge]: http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[license-link]: https://github.com/b4b4r07/go-crowi/blob/master/LICENSE

[godoc-badge]: http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
[godoc-link]: http://godoc.org/github.com/b4b4r07/go-crowi