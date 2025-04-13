package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	fmt.Println("Awaiting signal!")
	sig := <-sigs
	fmt.Printf("Got signal %s \n", sig.String())
}
