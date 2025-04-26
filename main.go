package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Args struct {
	Port      int
	Directory string
}

func Main() error {
	args := &Args{}
	flag.IntVar(&args.Port, "p", 8080, "Port")
	flag.StringVar(&args.Directory, "d", "", "Root of public directory")
	flag.Parse()

	err := http.ListenAndServe(fmt.Sprintf(":%d", args.Port), http.FileServer(http.Dir(args.Directory)))
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
