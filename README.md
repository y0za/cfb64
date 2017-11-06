# cfb64
convert file to/from Base64

## Usage
```
$ cfb64
NAME:
   cfb64 - convert file to/from Base64

USAGE:
   cfb64 [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
     encode, e  Encode file to Base64
     decode, d  Decode Base64 to file
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Example
It is useful when you want to send Base64 data with curl.
```console
$ cfb64 e --uri image.jpg | xargs printf '{"image": "%s"}' | curl -X POST --data @- http://localhost/api/images
```

## Installation
```
$ go get github.com/y0za/cfb64
```

## License
MIT License
