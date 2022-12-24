package main

import (
	"io"
	"log"
	"os"
	"os/signal"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		os.Exit(0)
	}()

	if err != nil {
		log.Fatal(err)
	}

	t := clipboard.FmtText
	b, _ := io.ReadAll(os.Stdin)
	clipboard.Write(t, b)

}
