go-crowi [![][travis-badge]][travis-link] [![][license-badge]][license-link] [![][godoc-badge]][godoc-link]
========

The official Go client for [Crowi](http://site.crowi.wiki/).

## Getting

go-crowi supports up to [v1.6.0](https://github.com/crowi/crowi/releases/tag/v1.6.0).

```console
$ go get github.com/crowi/go-crowi
```

Currently, go-crowi implements the following endpoints:

- `/_api/pages.create`
- `/_api/pages.update`
- `/_api/attachments.add`

These may be changed or newly added as Crowi changes.

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

[travis-badge]: http://img.shields.io/travis/crowi/go-crowi.svg?style=flat-square
[travis-link]: https://travis-ci.org/crowi/go-crowi

[license-badge]: http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[license-link]: https://github.com/crowi/go-crowi/blob/master/LICENSE

[godoc-badge]: http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
[godoc-link]: http://godoc.org/github.com/crowi/go-crowi
