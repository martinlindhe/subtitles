## About

Golang cli tool and library for reading, writing and manipulating .srt subtitle files


## Installation

Requires golang to be installed on your system.

```
go get github.com/martinlindhe/go-subber
```

## Usage

To download subtitles for a video file:

```
$ go-subber movie.mp4

Downloading subs for movie.mp4 ...
Cleaning sub movie.srt ...
Removing caption 82 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Removing caption 701 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Written to movie.srt
```

To specify another language than default "en", use the `--language="sv"` flag.


Remove ads from an existing .srt file:

```
$ go-subber subtitle.srt

Removing caption 21 [<font color="#FFFF00"> Captions by VITAC  </font><font color="#00FFFF"> www.vitac.com</font>]
Removing caption 22 [CAPTIONS PAID FOR BY DISCOVERY COMMUNICATIONS]
```

Strip html tags from .srt:

```
$ go-subber subtitle.srt --filter="html"

[html] <i>And there's a lot of it there.</i> -> And there's a lot of it there.
```

Normalize capitalization in .srt:

```
$ go-subber subtitle.srt --filter="caps"

[caps] INTRODUCING -> Introducing
[caps] right, to go? -> Right, to go?
```

A backup of the modified .srt file is created as .srt.org by default. To disable this behaviour, add the `--skip-backups` flag


## License

BSD


## Contributing

Patches welcome!

Some ideas, in no particular order:
- expose thesubdb.com search api
- .ssa/.ass reader/writer
- filter: spell fixer
- filter: remove hearing aid tags [DOOR OPENS]
- make -v (verbose mode) have any effect

- maybe .sub reader and converter to .srt
    - crappy frame-based format, but useful to be able to convert away from
    - requires specifying a frame rate
