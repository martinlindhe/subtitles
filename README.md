# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/subtitles?status.svg)](https://godoc.org/github.com/martinlindhe/subtitles)



Subtitles is a go library for handling .srt and .ssa subtitles

WARNING: The API is unstable, work in progress!


# Installation

```
go get -u github.com/martinlindhe/subtitles
```

# Example - convert srt to vtt

```go
in := "1\n" +
    "00:00:04,630 --> 00:00:06,018\n" +
    "Go ninja!\n" +
    "\n" +
    "1\n" +
    "00:01:09,630 --> 00:01:11,005\n" +
    "No ninja!\n"

res, _ := NewFromSRT(in)

// Output: WEBVTT
//
// 00:00:04.630 --> 00:00:06.018
// Go ninja!
//
// 00:01:09.630 --> 00:01:11.005
// No ninja!
fmt.Println(res.AsVTT())
```

# Example - download subtitle from thesubdb.com

```go
f, _ := os.Open(fileName)
defer f.Close()

finder := NewSubFinder(f, fileName, "en")

text, err := finder.TheSubDb()
```


# See also

- [subber](https://github.com/martinlindhe/subber) command line tool for subtitles
- [ssa2srt](https://github.com/martinlindhe/ssa2srt) for converting .ssa to .srt


# License

Under [MIT](LICENSE)
