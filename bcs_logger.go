package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	signalChannel = make(chan os.Signal, 1)
)

func main() {

	signal.Notify(signalChannel, syscall.SIGINT)
	signal.Notify(signalChannel, syscall.SIGTERM)

	go func() {
		for sig := range signalChannel {
			log.Fatal(fmt.Sprintf("signal=%d finished=%v", sig, time.Now()))
			os.Exit(1)
		}
	}()

	for true {
		fmt.Print(".")
		time.Sleep(1000 * time.Millisecond)
	}
}
