package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const cliVersion = "0.1.0"

const helpMessage = `
fileserver (v%s) runs a local server, which allows you to access files in the specified system directory.
	$ fileserver <path to directory>

By default, fileserver executes using port 8080 serving the files in current working directory.

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	port := flag.Int("port", 8080, "Server listening port.")
	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// Non flags. arguments
	args := flag.Args()

	if *version {
		fmt.Printf("fileserver v%s\n", cliVersion)
		os.Exit(0)
	} else if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Directory to serve
	path := "."
	if len(args) > 0 {
		fileInfo, err := os.Stat(args[0])
		if err != nil {
			fmt.Printf("Path %s is not a valid directory, an error was encountered.\n\t%s\n", args[0], err)
			os.Exit(1)
		}

		if !fileInfo.IsDir() {
			// is not a directory
			fmt.Printf("Path %s is not a directory, fileserver requires a valid directory.\n", args[0])
			os.Exit(1)
		}

		path = args[0]
	}

	fmt.Printf("Starting fileserver in http://localhost:%d for directory %s\n", *port, path)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), http.FileServer(http.Dir(path)))
		if err != nil {
			fmt.Printf("fileserver could not start a server in localhost:%d.\n%s\n", *port, err)
			os.Exit(1)
		}
	}()

	// Enabling shutdown with SIGINT (CTRL+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Block until we receive our signal.
	<-c
	fmt.Println("Exiting fileserver")
	os.Exit(0)

}
