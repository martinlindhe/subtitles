# About

[![Travis-CI](https://api.travis-ci.org/martinlindhe/subtitles.svg)](https://travis-ci.org/martinlindhe/subtitles)
[![GoDoc](https://godoc.org/github.com/martinlindhe/subtitles?status.svg)](https://godoc.org/github.com/martinlindhe/subtitles)

This is a go library and command-line tools for handling .srt, .vtt and .ssa subtitles

## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/subtitles/releases)

Or install them directly from git:

```sh
go install github.com/martinlindhe/subtitles/cmd/subber@latest

```

## Sub-projects

- [subber](https://github.com/martinlindhe/subtitles/tree/master/cmd/subber) command line tool for subtitles

## Library example - convert srt to vtt

```go
import "github.com/martinlindhe/subtitles"

in := "1\n" +
    "00:00:04,630 --> 00:00:06,018\n" +
    "Go ninja!\n" +
    "\n" +
    "1\n" +
    "00:01:09,630 --> 00:01:11,005\n" +
    "No ninja!\n"
res, _ := subtitles.NewFromSRT(in)

// Output: WEBVTT
//
// 00:04.630 --> 00:06.018
// Go ninja!
//
// 01:09.630 --> 01:11.005
// No ninja!
fmt.Println(res.AsVTT())
```

## Library example - download subtitles

```go
f, _ := os.Open(fileName)
defer f.Close()

finder := subtitles.NewSubFinder(f, fileName, "en")
text, err := finder.TheSubDb()
```

## License

Under [MIT](LICENSE)
