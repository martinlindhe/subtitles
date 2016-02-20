# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/subtitles?status.svg)](https://godoc.org/github.com/martinlindhe/subtitles)



Subtitles is a go library for handling .srt and .ssa subtitles

WARNING: The API is unstable, work in progress!


# Installation

```
go get github.com/martinlindhe/subtitles
```


# Example

Fetch subtitle from thesubdb.com:
```go
f, _ := os.Open(fileName)

finder := NewSubFinder(f, fileName, "en")

text, err := finder.TheSubDb()
```


# See also

- [subber](https://github.com/martinlindhe/subber) command line tool for subtitles
- [ssa2srt](https://github.com/martinlindhe/ssa2srt) for converting .ssa to .srt


# License

Under [MIT](LICENSE)
