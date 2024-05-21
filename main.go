package main

import (
	"github.com/witekblach/fridge/cmd/cli"
	"log/slog"
	"os"
)

func main() {
	err := cli.Execute(os.Args[1:])
	if err != nil {
		slog.Error(err.Error())
	}
}
