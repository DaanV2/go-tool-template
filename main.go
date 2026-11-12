package main

import (
	"context"
	"os"

	"github.com/DaanV2/go-tool-template/cmd"
)

func main() {
	if err := cmd.Execute(context.Background()); err != nil {
		os.Exit(1)
	}
}
