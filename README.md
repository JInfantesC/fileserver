# fileserver
fileserver runs a local server, which allows you to access files in the specified system directory.
```shell
$ fileserver <path to directory>

# Actions
$ fileserver -help

By default, fileserver executes using port 8080 serving the files in current working directory.

  -help
        Print help message and exit
  -port int
        Server listening port. (default 8080)
  -version
        Print version string and exit
```

As usual, you can build this using `go build main.go`. I've set a `Makefile` to just build and output in `dist/` using `make build`.