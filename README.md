## About

Golang cli tool and library for reading, writing and manipulating .srt subtitle files


```
go get github.com/martinlindhe/go-subber
```

## Cli Usage

To remove ads and clean up the formatting of a .srt file:

```
$ go-subber file.srt

Removing ads from seq 21 [<font color="#FFFF00"> Captions by VITAC  </font><font color="#00FFFF"> www.vitac.com</font>]
Removing ads from seq 22 [CAPTIONS PAID FOR BY DISCOVERY COMMUNICATIONS]
```


## License

BSD
