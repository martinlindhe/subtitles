## About

Golang cli tool and library for reading, writing and manipulating .srt subtitle files


```
go get github.com/martinlindhe/go-subber
```

## Usage

To download subs for a video file:

```
$ go-subber movie.mp4

Downloading subs for input file ...
Fetching http://api.thesubdb.com/?action=download&hash=691291516e111c67ba8d0246390d3bc1&language=en ...
Cleaning sub movie.srt ...
Removing caption 82 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Removing caption 701 [== sync, corrected by <font color="#00FF00">elderman</font> ==]
Written to movie.srt
```


To remove ads and clean up the formatting of an existing .srt file:

```
$ go-subber subtitle.srt

Removing caption 21 [<font color="#FFFF00"> Captions by VITAC  </font><font color="#00FFFF"> www.vitac.com</font>]
Removing caption 22 [CAPTIONS PAID FOR BY DISCOVERY COMMUNICATIONS]
```


## License

BSD


## Contributing

Patches welcome!

Ideas for features:
- cli flag to disable removing ads from subtitles (but still process file)
- cli flag to disable creation of .srt.org when cleaning up a .srt
- cli flag to specify sub language
- expose thesubdb.com search api
- .ssa/.ass reader/writer
- .sub reader and converter to .srt: requires specifying a frame rate
