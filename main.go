package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/ufukty/bs/internal/middlewares/with"
	"github.com/ufukty/bs/internal/middlewares/without"
)

func findport() (int, error) {
	for port := 8080; port < (1<<16)-1; port++ {
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			l.Close()
			return port, nil
		}
	}
	return -1, fmt.Errorf("no port available after 8080")
}

func Main() error {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	p, err := findport()
	if err != nil {
		return fmt.Errorf("finding available port: %w", err)
	}
	fmt.Printf("http://127.0.0.1:%d\n", p)

	h := with.Logging(without.Panic(http.FileServer(http.Dir(dir))))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), h); err != nil {
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
