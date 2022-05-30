package main

import (
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

func main() {
	var g errgroup.Group

	g.Go(func() error {
		return authRouter().ListenAndServe()
	})

	g.Go(func() error {
		return rasRouter().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
