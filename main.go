package main

import (
	"context"
	"github.com/evoaway/link-shortener/internal/server"
)

func main() {
	if err := server.Run(context.Background()); err != nil {
		panic(err)
	}
}
