## About

Subber is a cli tool and library for reading,
writing and manipulating .srt subtitle files written in Go


## Installation

Assuming golang is installed on your system:


For the subber tool:
```
go install github.com/martinlindhe/subber/subber
```

For the ssa2srt tool:
```
go install github.com/martinlindhe/subber/ssa2srt
```


## Usage

To download subtitles for a video file:

```
$ subber movie.mp4

Downloading subs for movie.mp4 ...
Cleaning sub movie.srt ...
Removing caption 82 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Removing caption 701 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Written to movie.srt
```

To specify another language than default "en", use the `--language="sv"` flag.


Remove ads from an existing .srt file:

```
$ subber subtitle.srt

Removing caption 21 [<font color="#FFFF00"> Captions by VITAC  </font><font color="#00FFFF"> www.vitac.com</font>]
Removing caption 22 [CAPTIONS PAID FOR BY DISCOVERY COMMUNICATIONS]
```

Strip html tags from .srt:

```
$ subber subtitle.srt --filter="html"

[html] <i>And there's a lot of it there.</i> -> And there's a lot of it there.
```

Normalize capitalization in .srt:

```
$ subber subtitle.srt --filter="caps"

[caps] INTRODUCING -> Introducing
[caps] right, to go? -> Right, to go?
```

A backup of the modified .srt file is created as .srt.org by default. To disable this behaviour, add the `--skip-backups` flag


Convert a .ssa to .srt:

```
$ ssa2srt subtile.ssa

[ads] 942 [Subtitles by: Scandinavian Text Service and JHS International ApS] matched subtitles by
Written subtile.srt
```

## License

BSD


## Contributing

Patches welcome!

Some ideas, in no particular order:

- automatically convert downloaded .ssa files to .srt, needs looksLikeSrt() and looksLikeSsa()
- expose thesubdb.com search api
- filter: spell fixer
- filter: remove hearing aid tags [DOOR OPENS]
- make -v (verbose mode) have any effect

- maybe .sub reader and converter to .srt
    - crappy frame-based format, but useful to be able to convert away from
    - requires specifying a frame rate
