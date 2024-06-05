## About

Subber is a cli tool for reading,
writing and manipulating .srt subtitle files


### Features

- Processes .srt and .ssa subtitles
- Cleans up the captions
- Optionally remove html tags or fix capitalization
- Outputs utf8 subtitle in the popular .srt format
- Synchronize captions


### Why?

While VLC may swallow all kinds of malformed subtitles,
not all software are so forgiving.
So, sometimes "other software" needs a cleaned up .srt,
and the `subber` cli app automates this task.


## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/subtitles/releases)

Or if you have golang installed, you can compile from source:
```
go install github.com/martinlindhe/subtitles/cmd/subber@latest
```


## Usage

### Download subtitles for a video

```
$ subber movie.mp4

Downloading subs for movie.mp4 ...
Cleaning sub movie.srt ...
Removing caption 82 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Written to movie.srt
```

Some additional flags (use `-h` for full list) includes:

  * `--keep-ads` do not remove advertising segments from the processesed sub
  * `--dont-touch` write sub as-is (downloaded are processed by default)`
  * `--skip-backups` by default, a .org file is created for every modified .srt
  * `--language="sv"` Specify another language, English subtitles is the default
  * `--sync="-3000"` Moves all captions backwards by 3000 ms


### Remove ads from an existing .srt file

```
$ subber subtitle.srt

Removing caption 21 [<font color="#FFFF00"> Captions by VITAC  </font><font color="#00FFFF"> www.vitac.com</font>]
```

### Strip html from .srt

```
$ subber subtitle.srt --filter="html"

[html] <i>And there's a lot of it there.</i> -> And there's a lot of it there.
```

### Normalize capitalization in .srt

```
$ subber subtitle.srt --filter="caps"

[caps] INTRODUCING -> Introducing
[caps] right, to go? -> Right, to go?
```

### Convert .ssa to .srt

```
$ subber subtitle.ssa -o subtitle.srt
```

## License

Under [MIT](LICENSE)
