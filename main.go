package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ufukty/bs/internal/middlewares/with"
	"github.com/ufukty/bs/internal/middlewares/without"
)

func Main() error {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8080),
		with.Logging(without.Panic(http.FileServer(http.Dir(dir)))),
	)
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
